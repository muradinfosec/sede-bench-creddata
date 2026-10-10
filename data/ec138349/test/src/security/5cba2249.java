/*
 * Copyright 2006-2011 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
 * the License. You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */
package org.springframework.security.jwt.crypto.cipher;

import java.math.BigInteger;

/**
 * @author Luke Taylor
 */
public class RsaTestKeyData {
	static final byte[] nBytes;

	static final byte[] dBytes;

	static {
		int[] nInts = new int[] { 161, 248, 22, 10, 226, 227, 201, 180, 101, 206, 141, 45, 101, 98, 99, 54, 43, 146,
				125, 190, 41, 225, 240, 36, 119, 252, 22, 37, 204, 144, 161, 54, 227, 139, 217, 52, 151, 197, 182, 234,
				99, 221, 119, 17, 230, 124, 116, 41, 249, 86, 176, 251, 138, 143, 8, 154, 220, 75, 105, 137, 60, 193,
				51, 63, 83, 237, 208, 25, 184, 119, 132, 37, 47, 236, 145, 79, 228, 133, 119, 105, 89, 75, 234, 66,
				128, 211, 44, 15, 85, 191, 98, 148, 79, 19, 3, 150, 188, 110, 155, 223, 110, 189, 210, 189, 163, 103,
				142, 236, 160, 198, 104, 247, 1, 179, 141, 191, 251, 56, 200, 52, 44, 226, 254, 109, 39, 250, 222, 74,
				90, 72, 116, 151, 157, 212, 185, 207, 154, 222, 196, 199, 91, 5, 133, 44, 44, 15, 94, 248, 165, 193,
				117, 3, 146, 249, 68, 232, 237, 100, 193, 16, 198, 182, 71, 96, 154, 164, 120, 58, 235, 156, 108, 154,
				215, 85, 49, 48, 80, 99, 139, 131, 102, 92, 111, 111, 122, 130, 163, 150, 112, 42, 31, 100, 27, 130,
				211, 235, 242, 57, 34, 25, 73, 31, 182, 134, 135, 44, 87, 22, 245, 10, 248, 53, 141, 154, 139, 157, 23,
				195, 64, 114, 143, 127, 135, 216, 154, 24, 216, 252, 171, 103, 173, 132, 89, 12, 46, 207, 117, 147, 57,
				54, 60, 7, 3, 77, 111, 96, 111, 158, 33, 224, 84, 86, 202, 229, 233, 161 };
		int[] dInts = new int[] { 18, 174, 113, 164, 105, 205, 10, 43, 195, 126, 82, 108, 69, 0, 87, 31, 29, 97, 117,
				29, 100, 233, 73, 112, 123, 98, 89, 15, 157, 11, 165, 124, 150, 60, 64, 30, 63, 207, 47, 44, 211, 189,
				236, 136, 229, 3, 191, 198, 67, 155, 11, 40, 200, 47, 125, 55, 151, 103, 31, 82, 19, 238, 216, 193, 90,
				37, 216, 213, 206, 160, 2, 94, 227, 171, 46, 139, 127, 121, 33, 111, 198, 59, 234, 86, 39, 83, 180, 6,
				68, 198, 161, 81, 39, 217, 178, 149, 69, 64, 160, 187, 225, 163, 5, 86, 152, 45, 78, 159, 222, 95, 100,
				37, 241, 77, 75, 113, 52, 65, 181, 93, 199, 59, 155, 74, 237, 204, 146, 172, 227, 146, 126, 55, 245,
				125, 12, 253, 94, 117, 129, 250, 81, 44, 143, 73, 97, 169, 235, 11, 128, 248, 168, 7, 70, 114, 138, 85,
				255, 70, 71, 31, 52, 37, 6, 59, 157, 83, 100, 47, 94, 222, 30, 132, 214, 19, 8, 26, 250, 92, 34, 208,
				81, 40, 91, 214, 59, 148, 59, 86, 93, 137, 138, 5, 104, 84, 19, 229, 60, 60, 108, 101, 37, 255, 31,
				227, 78, 61, 220, 112, 240, 213, 100, 80, 253, 164, 139, 161, 46, 16, 78, 157, 235, 159, 184, 24, 129,
				225, 196, 189, 242, 93, 146, 71, 244, 80, 200, 101, 146, 121, 104, 231, 115, 52, 244, 65, 79, 117, 167,
				80, 225, 57, 84, 110, 58, 138, 115, 157 };
		nBytes = new byte[nInts.length];
		dBytes = new byte[dInts.length];

		for (int i = 0; i < nInts.length; i++)
			nBytes[i] = (byte) nInts[i];
		for (int i = 0; i < dInts.length; i++)
			dBytes[i] = (byte) dInts[i];
	}

	public final static BigInteger N = new BigInteger(1, nBytes);

	public final static BigInteger E = new BigInteger(1, new byte[] { 1, 0, 1 });

	public final static BigInteger D = new BigInteger(1, dBytes);

	public final static String SSH_PRIVATE_KEY_STRING = "-----BEGIN RSA PRIVATE KEY-----\n"
			+ " MIIEowIBAAKCAQEAwARN4S7Z0asSEj61+SIvtUUuHopd/ffne1CbaHXNxW/TG4Xg\n "
			+ "9H7aHT7GaUOMcvj4pkSqzJzHJNWv+cykaFjdIPbI64G4vmw2q1dSVnQjcwksl+TG \n"
			+ "3KkGyRfjr940F6K1LIMjERZs9NxpgLKjmQwxIN1t5facX8atmNhUnyOMM4topbhU\n"
			+ "uf6EJUslRCjrMoNa2pW34kWDp5xlTczuOms6zKf49aNnmgxd8QHHVveWbj4SucPP\n"
			+ " hXgfrXkcVCcGkFTXx0wih5bmROmTunvmez+0urV0v5Pzfsu4OKbB9oe9je8SQCdh\n"
			+ "p1vZy6xbGvi5PEjJYH7nvw4Rbnwkr88SNgMZ1MnLlbLGZeaRDgRwN1N5oHriY/PY\n"
			+ "8xR3bVVyCFSchgFnBnPWRdetskZt8hyO5VqNZO9G58R7i8Ux9wbeE9bE4kc9Qgjd\n"
			+ " sxJoMfEORbDToInw4ka98dI278o43pBeM2CtGbDA3DpAEuJ21BUFiSljEBtdLQDw\n"
			+ " AnrRmJAUtGsOqfMVqTFhTK8wQKg7LzeoBALqxty7AdcLIpoddpkwmocAhJPJmSTK\n"
			+ "/RyMglS2HK1p1/s7uL+s9sesFkx2vT0jwXfBfkOROqS4j5wDbAStdut2FAhkOMLZ\n"
			+ "jnZC9Etzcqi6RNPCqh/PmeNwHiDBUAib9Uhd+pFPx7/anNrDVxRq83HFEPgxWGrG\n"
			+ "BujvGjalkTNp0lnH5UeDt33j+UZALJ4J5IML3q1iqpfbR9k5uwo6z8TE3UN6zTas\n"
			+ "tq0385OPZiiDXqetiC6YvDcHmgUXf9iU8Dd16Sj/U7VOLjlGeHSrv5rEyZJS6+3a\n"
			+ "1hdKrX82ZtsY2NeyoQ+EwyOg6hzQNQdNyHOYaRQg/crnGLP1YoMDsdzEXNyn9K40\n"
			+ "srCZiPw/x198Sm6gXh8MG1ZGtAsJ/Uh4pXUw8ipts2VTi8vwEEjbkzMeWAo8V2qZ\n"
			+ "VKOU/eIDOXb7B/CuIO4T/7+gPcSm2pk/1CeTAGONeWLdwVaJGJ6iVTvysgbX1iKS\n"
			+ "HgLcpFO4hBT7P6k3DHa4VERtTbDTVPUd61Gg0JAWxRj3NUxSB8sOPzjh3TSTMDOS\n"
			+ "kKol48BbzmZ2FBl95XV7FzsFYA/D3O6wgd/LIIG+mc7c/eC+3a1Q1CamSyY6NjH5\n"
			+ "HAVI31oQr83mhyiMRyCHgayGA2r7bu+VhbvdYMY2ldmpqWMRF5yola7qnp7PWEVq\n"
			+ "SlLNRiK5aUhT1BPX8bBkxlJKAevSvIWDYzOp4ri4vBvye9YQ7CYK8ycQnf64e+kg\n"
			+ "woyEaYciO3e39r+gGRf/OEDF1mjRIhD82DkY2kPyKS0YQZUCD/vX5b8+hDTLnJDQ\n"
			+ "cW6VDL9VeKHWzRcBhA7fvHcnxfdKUyV57aF5RbvqUiteFudPlck2QcAnGEXUMyaF\n"
			+ "QhXaQtaiQz8uzBmg1X/HPYbmvGEgeZ3EOo2cWbGmPQ8bD9ZtXFrolLlCMkbpDbFf\n"
			+ "lpRgdFYhmE5FpjRdZqXdHh4ibxZQ2ZbiN5OCb87YarGt58UnifO+Bj+PMsQrTyt6\n"
			+ " 6Ad5/16vsJno5VYCIOqf/yd/2FKNVpJtcWNARDgmv0ckcdb50f2T\n" + "-----END RSA PRIVATE KEY-----";

	public final static String SSH_PRIVATE_KEY_STRING_WITH_WHITESPACE = SSH_PRIVATE_KEY_STRING + "   ";

	public static final String SSH_PUBLIC_KEY_STRING = "ssh-rsa "
			+ "AAAAB3NzaC1yc2EAAAADAQABAAABAQDABE3hLtnRqxISPrX5Ii+1RS4eil399+d7UJtodc3GP9wjitjSTkQtnZIYIVWCbH1cAMkmFi"
			+ "hAiY768zEUCptbHI7jkHWtlHWrQKEQVYsY+Y0H59jxOsUciMlvzvezR3YVNgM7Cy3od0o+NeQt2AYc7grV8uc7VxAcoaCUEq9zhlJW"
			+ "eJQ4nUR0pkfEaeRWOmzWE/3wC9DHoSmYoTF7B3gwyPvuBFgB5KjSk+G6AuubLkMs+jFJQZJkQcI+UJ859MC3024ThjBukLAN8OZBv7"
			+ "2d6rtDQ/Ca0/qMWtXhVneKvZxZg5MXoVwvtkidwbdoK9fGnylRDs0+KZh3vR0Q+67V blah@blah.local";

	public static final String SSH_PUBLIC_KEY_STRING_WITHOUT_COMMENT = "ssh-rsa "
			+ "AAAAB3NzaC1yc2EAAAADAQABAAABAQDABE3hLtnRqxISPrX5Ii+1RS4eil399+d7UJtodc3GP9wjitjSTkQtnZIYIVWCbH1cAMkmFi"
			+ "hAiY768zEUCptbHI7jkHWtlHWrQKEQVYsY+Y0H59jxOsUciMlvzvezR3YVNgM7Cy3od0o+NeQt2AYc7grV8uc7VxAcoaCUEq9zhlJW"
			+ "eJQ4nUR0pkfEaeRWOmzWE/3wC9DHoSmYoTF7B3gwyPvuBFgB5KjSk+G6AuubLkMs+jFJQZJkQcI+UJ859MC3024ThjBukLAN8OZBv7"
			+ "2d6rtDQ/Ca0/qMWtXhVneKvZxZg5MXoVwvtkidwbdoK9fGnylRDs0+KZh3vR0Q+67V";


	public static final String SSH_PUBLIC_KEY_OPENSSL_PEM_STRING = "-----BEGIN PUBLIC KEY-----\n" +
			"MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwARN4S7Z0asSEj61+SIv\n" +
			"tUUuHopd/ffne1CbaHXNxj/cI4rY0k5ELZ2SGCFVgmx9XADJJhYoQImO+vMxFAqb\n" +
			"WxyO45B1rZR1q0ChEFWLGPmNB+fY8TrFHIjJb873s0d2FTYDOwst6HdKPjXkLdgG\n" +
			"HO4K1fLnO1cQHKGglBKvc4ZSVniUOJ1EdKZHxGnkVjps1hP98AvQx6EpmKExewd4\n" +
			"MMj77gRYAeSo0pPhugLrmy5DLPoxSUGSZEHCPlCfOfTAt9NuE4YwbpCwDfDmQb+9\n" +
			"neq7Q0PwmtP6jFrV4VZ3ir2cWYOTF6FcL7ZIncG3aCvXxp8pUQ7NPimYd70dEPuu\n" +
			"1QIDAQAB\n" +
			"-----END PUBLIC KEY-----";

	public static final String SSH_PUBLIC_KEY_PEM_STRING = "-----BEGIN RSA PUBLIC KEY-----\n" +
			"MIIBCgKCAQEAwARN4S7Z0asSEj61+SIvtUUuHopd/ffne1CbaHXNxj/cI4rY0k5E\n" +
			"LZ2SGCFVgmx9XADJJhYoQImO+vMxFAqbWxyO45B1rZR1q0ChEFWLGPmNB+fY8TrF\n" +
			"HIjJb873s0d2FTYDOwst6HdKPjXkLdgGHO4K1fLnO1cQHKGglBKvc4ZSVniUOJ1E\n" +
			"dKZHxGnkVjps1hP98AvQx6EpmKExewd4MMj77gRYAeSo0pPhugLrmy5DLPoxSUGS\n" +
			"ZEHCPlCfOfTAt9NuE4YwbpCwDfDmQb+9neq7Q0PwmtP6jFrV4VZ3ir2cWYOTF6Fc\n" +
			"L7ZIncG3aCvXxp8pUQ7NPimYd70dEPuu1QIDAQAB\n" +
			"-----END RSA PUBLIC KEY-----";

	public static final String SSH_X509_CERTIFICATE_PEM_STRING = "-----BEGIN CERTIFICATE-----\n" +
			"MIIDHDCCAgSgAwIBAgIJAK+wnYpjtdVFMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNV\n" +
			"BAMMGHNwcmluZy1zZWN1cml0eS1qd3QtdGVzdDAeFw0xODA0MTcwOTQ4MzVaFw0x\n" +
			"ODA1MTcwOTQ4MzVaMCMxITAfBgNVBAMMGHNwcmluZy1zZWN1cml0eS1qd3QtdGVz\n" +
			"dDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMAETeEu2dGrEhI+tfki\n" +
			"L7VFLh6KXf3353tQm2h1zcY/3COK2NJORC2dkhghVYJsfVwAySYWKECJjvrzMRQK\n" +
			"m1scjuOQda2UdatAoRBVixj5jQfn2PE6xRyIyW/O97NHdhU2AzsLLeh3Sj415C3Y\n" +
			"BhzuCtXy5ztXEByhoJQSr3OGUlZ4lDidRHSmR8Rp5FY6bNYT/fAL0MehKZihMXsH\n" +
			"eDDI++4EWAHkqNKT4boC65suQyz6MUlBkmRBwj5Qnzn0wLfTbhOGMG6QsA3w5kG/\n" +
			"vZ3qu0ND8JrT+oxa1eFWd4q9nFmDkxehXC+2SJ3Bt2gr18afKVEOzT4pmHe9HRD7\n" +
			"rtUCAwEAAaNTMFEwHQYDVR0OBBYEFPM7mHoBTz7Bgyblen9oSqd6gCVTMB8GA1Ud\n" +
			"IwQYMBaAFPM7mHoBTz7Bgyblen9oSqd6gCVTMA8GA1UdEwEB/wQFMAMBAf8wDQYJ\n" +
			"KoZIhvcNAQELBQADggEBAGfx6+D8YpYVHYbB9mdUDVmFKEq3rFBKaHXL8fDceHUi\n" +
			"GOAG0dLqP+lxx/pPsgfW8dnu1h/I5+cvOsj/YmwLMlodhrGN0XpaWmATz7+ikif3\n" +
			"VGGNXIWl/km+r30M4diFnSnycjYaOJdBqhLIkQd/w/JFFJ5J+C5b2281jYGw6Y1F\n" +
			"Kq3pqLlQVCnQhcnDroCtwLK78hG7yZasYVBnjKilSkMB1k14Kfq8WUR3NsODRiXg\n" +
			"EP+KsWrwS5l/cyUzkWDKgOvmlWeqSWp95WGhewuVAs34W0hzdT3JDd4TIX3NWMuw\n" +
			"i9txCbagsrq/2+rKgsasCPlcQwFw6Umzd73HuqiHmoM=\n" +
			"-----END CERTIFICATE-----\n";

}
