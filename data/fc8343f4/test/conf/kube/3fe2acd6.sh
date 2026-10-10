#!/usr/bin/env bash

# Copyright 2018 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

run_kubectl_config_set_tests() {
  set -o nounset
  set -o errexit

  create_and_use_new_namespace
  kube::log::status "Testing kubectl(v1:config set)"

  kubectl config set-cluster test-cluster --server="https://does-not-work"

  # Get the api cert and add a comment to avoid flag parsing problems
  cert_data=$(echo "#Comment" && cat "${TMPDIR:-/tmp}/apiserver.crt")

  kubectl config set clusters.test-cluster.certificate-authority-data "$cert_data" --set-raw-bytes
  r_written=$(kubectl config view --raw -o jsonpath='{.clusters[?(@.name == "test-cluster")].cluster.certificate-authority-data}')

  encoded=$(echo -n "$cert_data" | base64)
  kubectl config set clusters.test-cluster.certificate-authority-data "$encoded"
  e_written=$(kubectl config view --raw -o jsonpath='{.clusters[?(@.name == "test-cluster")].cluster.certificate-authority-data}')

  test "$e_written" == "$r_written"

  set +o nounset
  set +o errexit
}

run_kubectl_config_set_cluster_tests() {
  set -o nounset
  set -o errexit

  create_and_use_new_namespace
  kube::log::status "Testing kubectl config set-cluster"

  ca_file="${TMPDIR:-/tmp}/apiserver.crt"
  ca_data=$(base64 -w0 "$ca_file")

  # Set cert file
  kubectl config set-cluster test-cluster-1 --certificate-authority "$ca_file"
  expected="$ca_file"
  actual=$(kubectl config view --raw -o jsonpath='{.clusters[?(@.name == "test-cluster-1")].cluster.certificate-authority}')
  if [ "$expected" != "$actual" ]; then
    kube::log::error "Certificate authority did not match the expected value (expected=$expected, actual=$actual)"
    exit 1
  fi

  # Embed cert from file
  kubectl config set-cluster test-cluster-2 --embed-certs --certificate-authority "$ca_file"
  expected="$ca_data"
  actual=$(kubectl config view --raw -o jsonpath='{.clusters[?(@.name == "test-cluster-2")].cluster.certificate-authority-data}')
  if [ "$expected" != "$actual" ]; then
    kube::log::error "Certificate authority embedded from file did not match the expected value (expected=$expected, actual=$actual)"
    exit 1
  fi

  # Embed cert using process substitution
  kubectl config set-cluster test-cluster-3 --embed-certs --certificate-authority <(cat "$ca_file")
  expected="$ca_data"
  actual=$(kubectl config view --raw -o jsonpath='{.clusters[?(@.name == "test-cluster-3")].cluster.certificate-authority-data}')
  if [ "$expected" != "$actual" ]; then
    kube::log::error "Certificate authority embedded using process substitution did not match the expected value (expected=$expected, actual=$actual)"
    exit 1
  fi

  set +o nounset
  set +o errexit
}

run_kubectl_config_set_credentials_tests() {
  set -o nounset
  set -o errexit

  create_and_use_new_namespace
  kube::log::status "Testing kubectl config set-credentials"

  cert_file="${TMPDIR:-/tmp}/test-client-certificate.crt"
  key_file="${TMPDIR:-/tmp}/test-client-key.crt"

  cat << EOF > "$cert_file"
-----BEGIN CERTIFICATE-----
MIIDazCCAlOgAwIBAgIUdSrvuXs0Bft9Ao/AFnC7fNBqD+owDQYJKoZIhvcNAQEL
BQAwRTELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoM
GEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0yMDA1MTExODMyNDJaFw0yMTA1
MTExODMyNDJaMEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEw
HwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwggEiMA0GCSqGSIb3DQEB
AQUAA4IBDwAwggEKAoIBAQCrS5/kYM3TjdFw4OYMdSOys0+4aHPgtXqePWrn1pBP
W9D1yhjI/JS1/uPAWLi+1nnn0PBMCbqo+ZEsRIlSAABJE4fVbRPg+AuBPDdlCex5
QfKNi3vwp0Gy756SKTNZmIx42I9in0WeocCDliTEM2KsawCrVzuE3gwcHHkOnmLX
DEc4bajkQxcaJITG3hsJ2Cm5OBMLPwrsq/77VzOdC12r9j8+w0f7lCJfOm2ui7rm
Vl76V2Nits6U0ZrF1yzYtVQ1iWqCnOudPPf3jyc7KcSetGwozgoydkcqfUS9iMs9
2OV3v17GX6+sd8zY8tA95d/Vj6yU/l03GI9V6X9LXHSTAgMBAAGjUzBRMB0GA1Ud
DgQWBBQo2BKDxo4XI5FJDj9ZUuDst9ck7DAfBgNVHSMEGDAWgBQo2BKDxo4XI5FJ
Dj9ZUuDst9ck7DAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQAL
LjJ5k5kNdOUXFL2XhKO8w25tpkjd71vD3vKvPIN0Q5BgVmu5Bg476/WPBnSMvTVU
QYLQR5aGMLPdWuiSJnwO0BChKjOHaoCnP6D3cs2xP9hLBGoENaXMJRM4ZvFBRbES
Q6iTfq4OBPh5yCDDrjlppjhSnqaZuksmFnPFHLB/km003w8fgCD3VhmC2UFscl2K
nHaxK6uzIxisyE84ZDZYhjnPPib1wXGL8yu1dq0cbktE5+xJ2FHQkBJ6qaujkgV0
jpuWE9zk3CImFRkzPEwTF+3s5eP2XTIyWbtJGvJMmO0kHFx2PqCiAkdFldPKfrRh
M007Wf15dtkqyLNkzOxv
-----END CERTIFICATE-----
EOF

  cat << EOF > "$key_file"
-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCrS5/kYM3TjdFw
9jDoWnaOb2+8kTWxKEQyZzgG7uXWm8k9dcLR/az2/tlBVCM+4RUi9VVKVNtM+EBc
EjOgGKDqF9TTxKDH+pLcyYPGFcF9zUkSO0pdZ1Se706LlPvWqnK13p3La9hCfpmd
RpEfK8RxznWTVtOr8wyaphhBUQcPtkm2zIjcHKnPwYgU1Eub7Kr7BatcbHOGK/67
NERtH13r7n0+n9x0hcPDcT1gT6bjcF50S9crQp6r8gnL5OzTpCh0PQfeVABYYoB1
tPf7hNxXigZuauPLUDIdilR0zAt89ux9B21yZ3+YM4Gp1rl20J/dU3LZ/t34cx0S
3C0LPcNbWKxuEGtOpKCBsqMHoth+fAyJ1etjkutw2KXmOozbsE0bthPsCvbXspPu
iRmCJyvF/2sswcGy6vw1atZWqCgLJaRksRJOnDXv1lfnFdYSaXh36EhIJlPoUH7u
tUiimKopa4omwLdoBd5EUeRPg/O032t2P7DmyKoZiSV3FKMsdSHzeI42uZXxWWop
hx3nSoQOsqU5VXKSz1ioT4qAyj75fsfBsiUG0INktZF0L/SHjVGeh40SOrr4HfKQ
NhDbt+t7w3lVqrSJPZr+IOOhPigFD/naYCD8tOE7Fsq7nmxP5dKV0tGZDNU9OoMD
dTSrMpKcb6sjQicwTBykqL0bbg7OcKeF6cge3RXsyyBJRWXZz1BrJoJCnFCylN92
kCMtpfBFoygQG/PDuZZFdsElPV+RujycSyqnTibLMcNUSLFxpOIufj7MlS3UNLy5
ZXthWNNAJUCzCRAswYZ/eMcYCsBWHaNBkDppmGPA+QbcviWRbCihOkqTD4fRLciB
K0fMLfLxRsfY7ChAC7q1Ma8CzIdtLgTL6Mxi8s4m/I/4nY7z5xd7/yLK8lSVxX2U
ll8KX6fZ8XRm7ml/DUROxGmdhYwQB8mXkMyOr8QjpHCi9nmWUXfGUVZp29PhMvIY
M8ZBZuenHqm82BOPimHthupT3OjaWU5cemFJ+QNmZZ5599XpLqwahIPTjZAJQxy8
rJONd0EJvYNXeYts66e1fj029ZYvZgR2QabySXa2GOkttHDl5FfdGQt5/JPQP8qO
xxtQCVmgJzT1uQgWrS+QkRvHy8rq3jWXBiYgeTPN/a0DYtcRGGOzIJp0LI1XrRHc
BBnNMGWq1UcHPtB3S/hc4o8mYXI1TNarNlX+qFglH0p9FjU9f0ZciJFzXIxzfudm
KzUx4MBcWfchA0OdBekpSAGIHpDiFrx6o3TRiVokZ7PtiQ0tcQ/Big3+vm/gqBgh
dNjM+Amr1w6HFuyRhLsgEsUP/jUMVKMQNTU0fTIt8tbbQV82o1V1CPsByiTwe+Db
T+5Yq5wRHT8XMSnXBJGJW5viEXP5k46Fhiv3j5yCpSoULHyx4on+2XJpyH5VLmi5
Tw3IUkfGF+bHgNF5ZbrYNKXNAZ7FMs21I/L/NiT4bV1MZXjSR1ZM1afoJnFm0Gtl
PqZSavhi41yH7WG+LLZH9LdohNPzTdmbn39awUwP1TtPpvd7pk0bu67cDHYg+C6y
hWhQeBNKknYg03XUd+J74jY=
-----END PRIVATE KEY-----
EOF

  cert_data=$(base64 -w0 "$cert_file")
  key_data=$(base64 -w0 "$key_file")

  # Set client certificate and client key files
  kubectl config set-credentials user1 --client-certificate="$cert_file" --client-key="$key_file"
  expected="$cert_file"
  actual=$(kubectl config view --raw -o jsonpath='{.users[?(@.name == "user1")].user.client-certificate}')
  if [ "$expected" != "$actual" ]; then
    kube::log::error "Client certificate did not match the expected value (expected=$expected, actual=$actual)"
    exit 1
  fi
  expected="$key_file"
  actual=$(kubectl config view --raw -o jsonpath='{.users[?(@.name == "user1")].user.client-key}')
  if [ "$expected" != "$actual" ]; then
    kube::log::error "Client key did not match the expected value (expected=$expected, actual=$actual)"
    exit 1
  fi

  # Embed client certificate and client key from files
  kubectl config set-credentials user2 --client-certificate="$cert_file" --client-key="$key_file" --embed-certs=true
  expected="$cert_data"
  actual=$(kubectl config view --raw -o jsonpath='{.users[?(@.name == "user2")].user.client-certificate-data}')
  if [ "$expected" != "$actual" ]; then
    kube::log::error "Client certificate data embedded from file did not match the expected value (expected=$expected, actual=$actual)"
    exit 1
  fi
  expected="$key_data"
  actual=$(kubectl config view --raw -o jsonpath='{.users[?(@.name == "user2")].user.client-key-data}')
  if [ "$expected" != "$actual" ]; then
    kube::log::error "Client key data embedded from file did not match the expected value (expected=$expected, actual=$actual)"
    exit 1
  fi

  # Embed client certificate and client key using process substitution
  kubectl config set-credentials user3 --client-certificate=<(cat "$cert_file") --client-key=<(cat "$key_file") --embed-certs=true
  expected="$cert_data"
  actual=$(kubectl config view --raw -o jsonpath='{.users[?(@.name == "user3")].user.client-certificate-data}')
  if [ "$expected" != "$actual" ]; then
    kube::log::error "Client certificate data embedded using process substitution did not match the expected value (expected=$expected, actual=$actual)"
    exit 1
  fi
  expected="$key_data"
  actual=$(kubectl config view --raw -o jsonpath='{.users[?(@.name == "user3")].user.client-key-data}')
  if [ "$expected" != "$actual" ]; then
    kube::log::error "Client key data embedded using process substitution did not match the expected value (expected=$expected, actual=$actual)"
    exit 1
  fi

  set +o nounset
  set +o errexit
}

run_client_config_tests() {
  set -o nounset
  set -o errexit

  create_and_use_new_namespace
  kube::log::status "Testing client config"

  # Command
  # Pre-condition: kubeconfig "missing" is not a file or directory
  output_message=$(! kubectl get pod --context="" --kubeconfig=missing 2>&1)
  kube::test::if_has_string "${output_message}" "missing: no such file or directory"

  # Pre-condition: kubeconfig "missing" is not a file or directory
  # Command
  output_message=$(! kubectl get pod --user="" --kubeconfig=missing 2>&1)
  # Post-condition: --user contains a valid / empty value, missing config file returns error
  kube::test::if_has_string "${output_message}" "missing: no such file or directory"
  # Command
  output_message=$(! kubectl get pod --cluster="" --kubeconfig=missing 2>&1)
  # Post-condition: --cluster contains a "valid" value, missing config file returns error
  kube::test::if_has_string "${output_message}" "missing: no such file or directory"

  # Pre-condition: context "missing-context" does not exist
  # Command
  output_message=$(! kubectl get pod --context="missing-context" 2>&1)
  kube::test::if_has_string "${output_message}" 'context was not found for specified context: missing-context'
  # Post-condition: invalid or missing context returns error

  # Pre-condition: cluster "missing-cluster" does not exist
  # Command
  output_message=$(! kubectl get pod --cluster="missing-cluster" 2>&1)
  kube::test::if_has_string "${output_message}" 'no server found for cluster "missing-cluster"'
  # Post-condition: invalid or missing cluster returns error

  # Pre-condition: user "missing-user" does not exist
  # Command
  output_message=$(! kubectl get pod --user="missing-user" 2>&1)
  kube::test::if_has_string "${output_message}" 'auth info "missing-user" does not exist'
  # Post-condition: invalid or missing user returns error

  # test invalid config
  kubectl config view | sed -E "s/apiVersion: .*/apiVersion: v-1/g" > "${TMPDIR:-/tmp}"/newconfig.yaml
  output_message=$(! "${THIS_PLATFORM_BIN}/kubectl" get pods --context="" --user="" --kubeconfig="${TMPDIR:-/tmp}"/newconfig.yaml 2>&1)
  kube::test::if_has_string "${output_message}" "error loading config file"

  output_message=$(! kubectl get pod --kubeconfig=missing-config 2>&1)
  kube::test::if_has_string "${output_message}" 'no such file or directory'

  set +o nounset
  set +o errexit
}

run_kubeconfig_impersonate_tests() {
  set -o nounset
  set -o errexit

  kube::log::status "Testing config with impersonation"

  # copy the existing kubeconfig over and add a new user entry for the admin user impersonating a different user.
  kubectl config view --raw > "${TMPDIR:-/tmp}"/impersonateconfig.yaml
  cat << EOF >> "${TMPDIR:-/tmp}"/impersonateconfig.yaml
users:
- name: admin-as-userb
  user:
    # token defined in hack/testdata/auth-tokens.csv
    token: admin-token
    # impersonated user
    as: userb
    as-uid: abc123
    as-groups:
    - group2
    - group1
    as-user-extra:
      foo:
      - bar
      - baz
- name: as-uid-without-as
  user:
    # token defined in hack/testdata/auth-tokens.csv
    token: admin-token
    # impersonated uid
    as-uid: abc123
EOF

  kubectl create -f hack/testdata/csr.yml --kubeconfig "${TMPDIR:-/tmp}"/impersonateconfig.yaml --user admin-as-userb
  kube::test::get_object_assert 'csr/foo' '{{.spec.username}}' 'userb'
  kube::test::get_object_assert 'csr/foo' '{{.spec.uid}}' 'abc123'
  kube::test::get_object_assert 'csr/foo' '{{range .spec.groups}}{{.}} {{end}}' 'group2 group1 system:authenticated '
  kube::test::get_object_assert 'csr/foo' '{{len .spec.extra}}' '1'
  kube::test::get_object_assert 'csr/foo' '{{range .spec.extra.foo}}{{.}} {{end}}' 'bar baz '
  kubectl delete -f hack/testdata/csr.yml

  output_message=$(! kubectl get pods --kubeconfig "${TMPDIR:-/tmp}"/impersonateconfig.yaml --user as-uid-without-as 2>&1)
  kube::test::if_has_string "${output_message}" 'without impersonating a user'

  set +o nounset
  set +o errexit
}
