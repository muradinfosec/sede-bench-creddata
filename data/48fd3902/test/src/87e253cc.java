package org.cloudfoundry.credhub.views;

import java.io.IOException;
import java.time.Instant;
import java.util.UUID;

import com.fasterxml.jackson.databind.JsonNode;
import org.cloudfoundry.credhub.credential.RsaCredentialValue;
import org.cloudfoundry.credhub.domain.Encryptor;
import org.cloudfoundry.credhub.domain.RsaCredentialVersion;
import org.cloudfoundry.credhub.utils.JsonObjectMapper;
import org.cloudfoundry.credhub.utils.TestConstants;
import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.JUnit4;

import static org.cloudfoundry.credhub.helpers.JsonTestHelper.serializeToString;
import static org.hamcrest.CoreMatchers.equalTo;
import static org.junit.Assert.assertThat;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;


@RunWith(JUnit4.class)
public class RsaViewTest {

  private RsaCredentialVersion entity;
  private UUID uuid;
  private Encryptor encryptor;
  private Instant createdAt;
  private JsonNode metadata;

  @Before
  public void beforeEach() {
    encryptor = mock(Encryptor.class);
    uuid = UUID.randomUUID();
    RsaCredentialValue rsaValue = new RsaCredentialValue(TestConstants.RSA_PUBLIC_KEY_4096, TestConstants.PRIVATE_KEY_4096);
    entity = new RsaCredentialVersion(rsaValue, "/foo", encryptor);
    entity.setUuid(uuid);
    JsonObjectMapper objectMapper = new JsonObjectMapper();
    try {
      metadata = objectMapper.readTree("{\"name\":\"test\"}");
    } catch (IOException e) {
      e.printStackTrace();
    }
    entity.setMetadata(metadata);
    createdAt = Instant.now();
    entity.setVersionCreatedAt(createdAt);
    when(encryptor.decrypt(any()))
      .thenReturn(TestConstants.PRIVATE_KEY_4096);
  }

  @Test
  public void itCanCreateViewFromEntity() throws IOException {
    final Instant createdAtWithoutMillis = Instant.ofEpochSecond(createdAt.getEpochSecond());
    final RsaView actual = (RsaView) RsaView.fromEntity(entity);
    assertThat(serializeToString(actual), equalTo("{"
      + "\"type\":\"rsa\","
      + "\"version_created_at\":\"" + createdAtWithoutMillis + "\","
      + "\"id\":\"" + uuid + "\","
      + "\"name\":\"/foo\","
      + "\"metadata\":{\"name\":\"test\"},"
      + "\"value\":{"
      + "\"public_key\":\"-----BEGIN PUBLIC KEY-----\\n"
      + "MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA7PJUCHSfrZ3PY7n1/cC8\\n"
      + "wj1UtbraEycM0DtjUuRdOzhFl50feF7WGNyrQS7tdCx42cx+ZmsXVmJMm+BoGEsF\\n"
      + "fupZ/k8Z+/S0E/IErWyvpiQRVpxqZRzVPpaB3R5CnJDlYoQYA347FufLOOl/R5ig\\n"
      + "IlHrVwWs0F9qVToTigRE4BLLVSgdQqDdHz/fJm3q4Fvkxx8Q6W7aHXqdRKFmXj5/\\n"
      + "0Hv45cu46C/B2blQMM7p5gbK65tAdXERd5huP12Q3R6m89YbpqM//FfCjAbgvPfV\\n"
      + "4Tdcmb5nWantzoaRDa/Dt3tuqe8cU8gsjrOZJNjMr+2Kljsq+N6GxSMxlHCyzI96\\n"
      + "dUg7pIM6cQhz5ome9szv4Wfh5Aet9fYmT7GYK0fL2qULCWOTH2CTvSj49ASpRXYv\\n"
      + "cue/FQo9acOjmM+37ka0n7v1tEtZfRSyOOdGhbDUo4uWV+3pSeQNKg0GhW5RAgLJ\\n"
      + "mzI0/TAg8iQ42X5rW/VGOI/8sBXmD9ukS6m3VjGQWILpFYPorAssSXM3nGiSkVgZ\\n"
      + "1K8s+bEr27Dgt/K3buywJf78X/JtmWOnf02dkoExA4A+GMThulmRJts9iRqYBrUH\\n"
      + "FTuLpX0mv8aqL74nS3P5E1kdeXVbbPq9VL+wbWmse16kyqgZK8Y4W6Vw9kVUAHmY\\n"
      + "QTkPNKT2xbb1DzdaQHYHNeMCAwEAAQ==\\n"
      + "-----END PUBLIC KEY-----\","
      + "\"private_key\":"
      + "\"-----BEGIN RSA PRIVATE KEY----- fake\\n"
      + "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\\n"
      + "-----END RSA PRIVATE KEY-----\""
      + "}}"));
  }

  @Test
  public void hasVersionCreatedAtInTheView() {
    final Instant now = Instant.now();
    entity.setVersionCreatedAt(now);

    final RsaView actual = (RsaView) RsaView.fromEntity(entity);

    assertThat(actual.getVersionCreatedAt(), equalTo(now));
  }

  @Test
  public void hasTypeInTheView() {
    final RsaView actual = (RsaView) RsaView.fromEntity(entity);

    assertThat(actual.getType(), equalTo("rsa"));
  }

  @Test
  public void hasAUUIDInTheView() {
    final RsaView actual = (RsaView) RsaView.fromEntity(entity);

    assertThat(actual.getUuid(), equalTo(uuid.toString()));
  }

  @Test
  public void hasMetadataInTheView() {
    final RsaView actual = (RsaView) RsaView.fromEntity(entity);

    assertThat(actual.getMetadata(), equalTo(metadata));
  }
}
