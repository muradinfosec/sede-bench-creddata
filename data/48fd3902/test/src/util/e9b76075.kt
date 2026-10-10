package org.cloudfoundry.credhub.utils

class TestConstants {
    private constructor()

    companion object {
        const val INVALID_PRIVATE_KEY_NO_HEADERS: String = "some-invalid-private-key"

        const val INVALID_PRIVATE_KEY_WITH_HEADERS: String = "-----BEGIN RSA PRIVATE KEY----- fake\n" +
            "MIIEpQIBAAKCAQEAwqIrV8HpCuPyuJ6VvyG7gVhYJGAOX4zhclxuYFbR9Rni5Nrh\n" +
            "866zhZrBTA+tHT+rGzuAGtLQUBrqeUIf2qK2pL0MTzNDh95VtLhLSf66iagBFwkx\n" +
            "HbwYP/5AoBUsCd5oGyZ29pdRLRoyfNW0yMjb/03Qdm9TNmUKBn0Bi5d1pIXCSTVf\n" +
            "ByonkRtXBQyFiFhYw3ABHQQZz33nwNgY63QSZmOkZTYqKt9Brme9w1PoW1Y3Bkw8\n" +
            "Wir5qCMBFmnhOa2alooe8rNxXgqfpqvnoDrPBhxcC8IzWgt2rUgxOu6s924VDD8F\n" +
            "c6JwUCIZzn9fezKFVhdXcVqBoCCui8HhYjB3MjEWcRjWAmrMutexYbgrfNGcHWkd\n" +
            "9sPnxWgY9ImaTrHbS3adc8UWeuwdaA7L8MdyX6pceVBZDRHUv8hbICiwSceJs2S/\n" +
            "vdA/oVjgmRRHEu/Bzvql1OSjxKTIrTskAZZHz3W0sPoxKBP7XTcZVI/I3li14S8D\n" +
            "lkTYWiKCy33TCOVlVG5NdCJkkH/l4/cmGqPBj9dRIWJvOk0W6EnlE14jYLxz4apC\n" +
            "30SKsaiSXVVvBFqVQL7TRKevoAWx8ijDHQKnvsD9MiK5StgqOnAYfs/kxRfTh3NU\n" +
            "kBZy0Bhj2Gcvm/oYDKw7+ZCCFPRVrnPhmBCcLvvAGwYGnyTrxdBB7fu1JcxV11Za\n" +
            "cz6CuTKXLxWP/pDx40UMjGhHG13aSo/JcKM3XCdX7P1guSlprXBEzzQpATsze4CF\n" +
            "j6gEsuFAL609YSTnCfpxz6PWpzqFu0zwfwW+AF+ltAY3Tvv2pixBGpeGkW5tdA+x\n" +
            "8a5/TF/r3HDJrVH+f1SLJMaEKDtOn7JtMqsJGsw1bovB3WKNRct0kAZaHWjYU1PA\n" +
            "qgAFkie9lAoVXKUgtMFoONrF5CjrlmRDSWYbkn9UZcF82bRcKlT2hIy4kBg/Spvg\n" +
            "1gaXhdo8vgL//4PfGxYxB/mE+8vPYtphce/7VU675OVJBQxd4YiA1JkjjiWvCdTB\n" +
            "l4WWsHaEhU9xRq/b777Kk99CZF1QLYput6LgYMIAGxPe0F6mhTDV44AY+DUBLuhC\n" +
            "JyAFa5Ug/2Su2l1cfjLoRU3Xc6W9Vok+MgRqIHQ0aiNR9bZ0lDPymZAFxy9JjARn\n" +
            "yvBidh2iFrrxR6fi6sQh4uJMHrApWypOOxrY1EplWGplJa6u8pqx68dtpYWN/cdf\n" +
            "RK0tAMd0k/UMqBfi+6c1FuQIkFndVmiXsnHJxbP9oBfIF7gUR664x0zRmg1YJJl9\n" +
            "IUXqRy2tkQ04xucAp2yXPEibKMfml75HmcL5nduPBIsrqjAX4qygwLfQnWAX+PBu\n" +
            "plr/QARZV61LHudW+B6VPDdJ6hX2fhWhF4DbGpquwonWiFJwlz2ObTjEojERhBUV\n" +
            "++3BuSqKUYdZpdE3eTSEm4RrQmZZa3wtx8Nmn0BDxS+N7cRansv9ew/ucPZgHPS4\n" +
            "KqGZWwKVL2xDFEe5Tv3j1PHY4DETyC0pNKVmLtC+AjHQ82FRlRllrS8OTixoWfw8\n" +
            "2bFQT1beRLxjpKSU5bPZ93QhF5WUTYwO1PSgAg3+QYLDiAqG4e4/byte=\n" +
            "-----END RSA PRIVATE KEY-----"

        // generate with ./build/credhub n -t ssh -n foo -k 4096
        const val SSH_PUBLIC_KEY_4096: String = (
            "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDs8lQIdJ+tnc9jufX9wLzCPVS1utoTJ" +
                "wzQO2NS5F07OEWXnR94XtYY3KtBLu10LHjZzH5maxdWYkyb4GgYSwV+6ln+Txn7" +
                "9LQT8gStbK+mJBFWnGplHNU+loHdHkKckOVihBgDfjsW58s46X9HmKAiUetXBaz" +
                "QX2pVOhOKBETgEstVKB1CoN0fP98mbergW+THHxDpbtodep1EoWZePn/Qe/jly7" +
                "joL8HZuVAwzunmBsrrm0B1cRF3mG4/XZDdHqbz1humoz/8V8KMBuC899XhN1yZv" +
                "mdZqe3OhpENr8O3e26p7xxTyCyOs5kk2Myv7YqWOyr43obFIzGUcLLMj3p1SDuk" +
                "gzpxCHPmiZ72zO/hZ+HkB6319iZPsZgrR8vapQsJY5MfYJO9KPj0BKlFdi9y578" +
                "VCj1pw6OYz7fuRrSfu/W0S1l9FLI450aFsNSji5ZX7elJ5A0qDQaFblECAsmbMj" +
                "T9MCDyJDjZfmtb9UY4j/ywFeYP26RLqbdWMZBYgukVg+isCyxJczecaJKRWBnUr" +
                "yz5sSvbsOC38rdu7LAl/vxf8m2ZY6d/TZ2SgTEDgD4YxOG6WZEm2z2JGpgGtQcV" +
                "O4ulfSa/xqovvidLc/kTWR15dVts+r1Uv7Btaax7XqTKqBkrxjhbpXD2RVQAeZh" +
                "BOQ80pPbFtvUPN1pAdgc14w=="
            )

        const val SSH_PUBLIC_KEY_4096_WITH_COMMENT: String = (
            "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDs8lQIdJ+tnc9jufX9wLzCPVS1utoTJ" +
                "wzQO2NS5F07OEWXnR94XtYY3KtBLu10LHjZzH5maxdWYkyb4GgYSwV+6ln+Txn7" +
                "9LQT8gStbK+mJBFWnGplHNU+loHdHkKckOVihBgDfjsW58s46X9HmKAiUetXBaz" +
                "QX2pVOhOKBETgEstVKB1CoN0fP98mbergW+THHxDpbtodep1EoWZePn/Qe/jly7" +
                "joL8HZuVAwzunmBsrrm0B1cRF3mG4/XZDdHqbz1humoz/8V8KMBuC899XhN1yZv" +
                "mdZqe3OhpENr8O3e26p7xxTyCyOs5kk2Myv7YqWOyr43obFIzGUcLLMj3p1SDuk" +
                "gzpxCHPmiZ72zO/hZ+HkB6319iZPsZgrR8vapQsJY5MfYJO9KPj0BKlFdi9y578" +
                "VCj1pw6OYz7fuRrSfu/W0S1l9FLI450aFsNSji5ZX7elJ5A0qDQaFblECAsmbMj" +
                "T9MCDyJDjZfmtb9UY4j/ywFeYP26RLqbdWMZBYgukVg+isCyxJczecaJKRWBnUr" +
                "yz5sSvbsOC38rdu7LAl/vxf8m2ZY6d/TZ2SgTEDgD4YxOG6WZEm2z2JGpgGtQcV" +
                "O4ulfSa/xqovvidLc/kTWR15dVts+r1Uv7Btaax7XqTKqBkrxjhbpXD2RVQAeZh" +
                "BOQ80pPbFtvUPN1pAdgc14w== dan@foo"
            )

        const
        val RSA_PUBLIC_KEY_4096: String = (
            "-----BEGIN PUBLIC KEY-----\n" +
                "MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA7PJUCHSfrZ3PY7n1/cC8\n" +
                "wj1UtbraEycM0DtjUuRdOzhFl50feF7WGNyrQS7tdCx42cx+ZmsXVmJMm+BoGEsF\n" +
                "fupZ/k8Z+/S0E/IErWyvpiQRVpxqZRzVPpaB3R5CnJDlYoQYA347FufLOOl/R5ig\n" +
                "IlHrVwWs0F9qVToTigRE4BLLVSgdQqDdHz/fJm3q4Fvkxx8Q6W7aHXqdRKFmXj5/\n" +
                "0Hv45cu46C/B2blQMM7p5gbK65tAdXERd5huP12Q3R6m89YbpqM//FfCjAbgvPfV\n" +
                "4Tdcmb5nWantzoaRDa/Dt3tuqe8cU8gsjrOZJNjMr+2Kljsq+N6GxSMxlHCyzI96\n" +
                "dUg7pIM6cQhz5ome9szv4Wfh5Aet9fYmT7GYK0fL2qULCWOTH2CTvSj49ASpRXYv\n" +
                "cue/FQo9acOjmM+37ka0n7v1tEtZfRSyOOdGhbDUo4uWV+3pSeQNKg0GhW5RAgLJ\n" +
                "mzI0/TAg8iQ42X5rW/VGOI/8sBXmD9ukS6m3VjGQWILpFYPorAssSXM3nGiSkVgZ\n" +
                "1K8s+bEr27Dgt/K3buywJf78X/JtmWOnf02dkoExA4A+GMThulmRJts9iRqYBrUH\n" +
                "FTuLpX0mv8aqL74nS3P5E1kdeXVbbPq9VL+wbWmse16kyqgZK8Y4W6Vw9kVUAHmY\n" +
                "QTkPNKT2xbb1DzdaQHYHNeMCAwEAAQ==\n" +
                "-----END PUBLIC KEY-----"
            )

        const val PRIVATE_KEY_4096: String = (
            "-----BEGIN RSA PRIVATE KEY----- fake\n" +
                "MIIJKAIBAAKCAgEA7PJUCHSfrZ3PY7n1/cC8wj1UtbraEycI3sVMYihUeYAWR98Q\n" +
                "UU0cXXexxZ5SssQ38Fd+UEBXRVoDf+NjQxvKiHMf/I1O+/r4f/qCdmziVkKQwmoY\n" +
                "stzylTXb0x4kTXGiYIeRL269aUKrWmk/m5ePyeNzRLUH8a6UPDTqrnRP0OfqVNeS\n" +
                "uUbggT/jGz8i0wRekq9c7T4vbLnEICjBdg1/9Be18ds67f/q3TpDjn8k7dsy89ag\n" +
                "tpMwX3dyQ21h8f4a22cuHlh//FizmozohGUk8uuhhh1pmDFwvZrFza/qV4kqHn3c\n" +
                "i4ViUFzEQQDLh+3lLKWE+b6yEELEardIoR14jtp0USV8FPin5SwN6pmI6eHU2qiM\n" +
                "9Xdxc6pvd9Zs9atzNOvcq0gdWHc27RNiWPRBUIA/yBp3hlRbAL+98at5u1x9qcVE\n" +
                "IgrLMxJlXYqFw1Axm+3UrcgVxS6zac6MhTfQigi7/LRK6mL95g4Kw/otUs/8ttXd\n" +
                "B6dUi4X0jVPfpsuhsKSXSuKxnQQ7OWazstVl8d6r+cYF56LDv/P7DTkIog73Z/HM\n" +
                "IuVJJ50vexyrw5t+SOMZXYnlmQz2mZFfNZnjCLmlCd5ue3xwQ83SW3p6r9VwcYHp\n" +
                "JUw9zj+ahVUmk37atmArV8Y4N3XN8TNcFrJzdQiZCMo0zGB0VzLcaPyHEYzdzxaM\n" +
                "ewyRJamyahuDaDNJSWL1RpGqgoAkniR0iCw8LdcN5gSTTqlgVRslT2edGP46Do4G\n" +
                "HRCF1i9cJHs++SC4E/SDylB0vR33iNpnVImVYt7icMq41dOYyFwpa9xTG5sxC4Qf\n" +
                "B6dikYlEAVGyKWmoAjBkxJD6pw/8MuJZH0ZO7IyAkeiTBHjEWW1nO/t5rGjDaeE/\n" +
                "GvyQZ+TINdCENqfi1sprZVjjJuIZ31CkNOG9IL9wG0bSK/XtezsbAln8II+LZy2K\n" +
                "DtxPgsw5Ll8yeAKADic1Z8seBqurVtay7fBNEoIhUaazAy9uDAwdmDG30g+Rt1F7\n" +
                "uCUrP0we6FRqSEj6hRWFMhRQHR6rftDfPb0Ey5SIGM6KuMwnEJ8mzNXZs/pPWjRf\n" +
                "4EWUO6764MVBQeBxAPP8bSMZBivzjm92ch6E51HqXtiP1WrXa4K0TLtC+kH8Sj3L\n" +
                "LMy3mb6ngoVqeP8x7TmwC6CmsfrFL/oaRX5UPw/AqU+pJFUjRTSviXG4BUT3FdZs\n" +
                "JmfXhh0R5sIsgc2iZPCMoNYA9XkAl2aT+uAAxthqjBDXJhMTkfy+S4z0GV0QrV4L\n" +
                "UausN9HHtlFT0GxuLhPpNpyrOFC8G2kvG0qSaxEwNJal4KGFTlTn9C4W3gLIGROo\n" +
                "fTDcBxt+YPQBIZ68mCCaJkJnWY1AvtA3ZCV0ZPyimoVc2fCvMjbJIGZr3rZqdjPu\n" +
                "DaAvYznqj30JUWx/CzKxZpSV9EidyvnCXyZBppaLHgFGZZAEpYz/tjbrdT9lODkI\n" +
                "RXqjJu/o2WEPIrvX4JsvXErNcazwkudICS9ycbZKnf9rsYrcjTinF9uwDKLwEywd\n" +
                "RAyZ8efOEwGJbzWCJPd0nXAIAcufqKwrvcyXyb5ZbhI6IptEXofGE1QPbPI5VuQH\n" +
                "IctJRiKwxkobmMePOh5bbT6HTJ5E3+pg94jV1b5I5e6uxGsWBWqrZQvybOfFSKZs\n" +
                "2cAv7YduCN1amsXFz7yyIsr/1iVH566E0GdiMDiqdz5Be5mMueBYpcFhcMPuIXkR\n" +
                "OHH/tHZTNV083bjkuffJ7T6Td5ZF5uu99kMB5wqHUEY0V+R7wdEiy1ZrYRt/d/FV\n" +
                "lqSak3M9TQg5qiF3ED4F119R4dg5eNaNg/IcDur8O0mPjnQFLoCz7Fa1jFhRCYaj\n" +
                "m9+pvCW2oiU3D5E00lnAf4ttdN/IeSqSBQDLIeVmzWNPfG41jYU31d1H3gqngym0\n" +
                "tBvwI/a4+SooYfRzxDuU3AwKuxo1o+W1FbGd4NEp+Yzw9f/fqF5PEtnVneMeOD7d\n" +
                "MTNvURVdajU5WwwiP6MXMnWWDhWtyFRahTG7pCk8exJ1v48wAnvDzn/wC2D5P6p6\n" +
                "89KRK7hyZ8gB80gG6lJFeoO/dTVlr8vUfvBPdcgGaDHvFxNjxmwQ7p6y/+LS/djr\n" +
                "rDfGl8zjVrRHVrUq9Ggkj3qRgST2nbuu8iPq2KMw5gp4vo9dBbEizW5vHYheUEO/\n" +
                "YmHECBI8Vgb5IPhkNlhWTvdaPxE8vwVTbGqN+wYVa/4J2XNMHHONFMc6YVdxdR13\n" +
                "QVoLIEGufIxZTbPwRNafNF/5BV9yOY9UsQ4wwRm0fJisQ1qeO+zugk7s5dEvZkqb\n" +
                "TS0cA5E1tDwS7NbA89GBwrdJLZJ1jBBNWqh8YrdQAeTEElFU2b3cl4ckhCbLAmg+\n" +
                "WrHcTFXDpwoRrsx3E6prdvTEsOVGET14NPS7zVymcd8vWSnlsIwSoR9aPG5WZFlw\n" +
                "QepjHSxwDexKAw4uG1EVAGVG+Qyx4PhO2ex5AKZ06J/eY1gtDhPgeCgIBGE7ceUL\n" +
                "at6RP9ygWvdhY1pA5R211D7T9q39xuTOsnboyx6dmm7aEOYQ4osi+maEcJ1Hgvcj\n" +
                "ixEJ3soQJOvkwoixJv79UQUIxJi9pDZibK4HvLA8Pt8JqW4XXnpXmAr4S7xzUTlM\n" +
                "yv9mnquCvTR7PZElICFKBQOmlp7jAGajOEqXyMstFlNoPiTTs5+D7PMEbUM2g3WF\n" +
                "YBLvtyXe5EUDa6Xg/ZmmRhRdDYh70FHoBfOB8qH2aZGEXKsx8Qa+E45uEFtkbUC8\n" +
                "y3rF8LzxsnNdQU9+DZTt69vmR7PwbQAngZnMbArHIVoeHcKdQcf0+AOoQMnOTYeh\n" +
                "5kDQTBI8CSzJqI2oQ2wTjOBLqolT2qRVbPJ3tBT0LbLrKF7jbAFKWUg0Hf7+3U+l\n" +
                "6Ujlq0JHW0ir+YPBQc9w9VREar3WeYRaU3eyGk3VyDEUuhe1zTy2NVLbSPxcF6TO\n" +
                "q+Q3i0sFe8yH2TNmP15huHeRE9vPDDF0o1oL02G6t30fjT8+VKF1ULig2ThfWitI\n" +
                "Gzj+QdwVzdSoJgdQwf+eJQRK+RTHVDgBdYFaWzOrMzZZc/c3rYcr1MeFqVpwN1uq\n" +
                "HTbZ9l/kdLdjc9qpF8QNXMoGk5UrqoIRQWam/YO8Vph3zt8uQhZAswO4K9P=\n" +
                "-----END RSA PRIVATE KEY-----"
            )

        const val TEST_CA: String = (
            "-----BEGIN CERTIFICATE-----\n" +
                "MIIC/zCCAeegAwIBAgIUbW0U9/70rhPUj7qsG+kplvNKxfYwDQYJKoZIhvcNAQEL\n" +
                "BQAwDzENMAsGA1UEAxMEdGVzdDAeFw0xOTA5MDQxODI5NTRaFw0yMDA5MDMxODI5\n" +
                "NTRaMA8xDTALBgNVBAMTBHRlc3QwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEK\n" +
                "AoIBAQC43kce29qg4h46x4b6K7WVxR4EQTKaa9JS+sj8NN8FJRQJljxNDiWH/0kM\n" +
                "YwG2VYq/XGibDDaEiGqhJ4ozbWiZcwczddPLQlRuV7XIwsY0RW/iqt2CT5hf6QXC\n" +
                "jQX3XD3hcez0S1KpEcYMycLYkoNARGxcQz1mh7EuyzCLN+Uh47JK56eOlUN4iS2o\n" +
                "xz9HfB5NNORBZi8QSPZdZ6VrKOzaeQw7bO4jptyHJSmFGU9FiOa2OTj/kYnWggJH\n" +
                "gmSE4qh5e55Lgl4aCr/Cl31DIChxnIgMz0qiCyuV111sBnKq3kyFrEejMHshIYeU\n" +
                "kk+iAGoOCaIsxkbTkL8id+8uszJxAgMBAAGjUzBRMB0GA1UdDgQWBBRio//8/deq\n" +
                "M5Bw98SgEqfAy4btJjAfBgNVHSMEGDAWgBRio//8/deqM5Bw98SgEqfAy4btJjAP\n" +
                "BgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQCEHLDlselE4BZ+angg\n" +
                "00eHEJ2RU4BaIdfzYgcCwj7rfh2HeHFl4OphM3XWjzDFuEDFc7u2dhfaSZbkHJtc\n" +
                "zGCkKPRkYMHci6Aq2nbeIXahKRHwfycQWbOKE4aOooblY7dH8iTbzrHQY7TO3sJV\n" +
                "owt82sKm+FMdjWI0CkCsK55X+zT8WkNPxo6xLWSppaRcGYLYPFGuBcIVWuFyqCN/\n" +
                "T4DSBV+SDFTp0SOd9R4HydAmwksar+tdhJoCW+N1WkF02q83NLkVX7rlPrJ9DPy7\n" +
                "EpKeflegYY8nMLDjhH7m8p5niNO0caiNEzHrpTs5cqlFJZQF15Lu9u+6/sYWsRAW\n" +
                "IH3c\n" +
                "-----END CERTIFICATE-----"
            )

        const val TEST_TRUSTED_CA: String = "-----BEGIN CERTIFICATE-----\n" +
            "MIIDTTCCAjWgAwIBAgIULxxoB3zfye0MzzRQGtKtw8CC2p4wDQYJKoZIhvcNAQEL\n" +
            "BQAwGjEYMBYGA1UEAwwPZm9vX2NlcnRpZmljYXRlMB4XDTE3MTEyMTE2MjQ1NFoX\n" +
            "DTE4MTEyMTE2MjQ1NFowGjEYMBYGA1UEAwwPZm9vX2NlcnRpZmljYXRlMIIBIjAN\n" +
            "BgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvgmVFC08CaGlp2ENKc5mym9BkcEL\n" +
            "E2030VXTJJSiLH1Py3s79rbJL4F/loxAMSbQCuKIADZZ4Wu7Xp8bFY92u0nNrAuG\n" +
            "Xw1omhfF5UTi1vewRG8inxdJZs9vxsiXZWI6WzUpvJZeCMvvKlkhA/C+GwjpIdRg\n" +
            "3fVZ44JvdZOx4cDUagRkDRcsHABp/ip19xhfZGLwFuJw1wd5kxZmZKCreoel+b/5\n" +
            "BaQyFD0L9OkSfq1Y/nMqOdIbXEnpeg8sawhQan0s7G98MTsvDF14jnucCECLO1fg\n" +
            "1YpxoBTDkNlrIPq4G8UO4+GNz5FJEIBGOsiRmEn0VjFEpZ3k+t/Nkf/b6wIDAQAB\n" +
            "o4GKMIGHMB0GA1UdDgQWBBQ3ZlJJaG9Brzf3IM6tWsMJce6YIDBVBgNVHSMETjBM\n" +
            "gBQ3ZlJJaG9Brzf3IM6tWsMJce6YIKEepBwwGjEYMBYGA1UEAwwPZm9vX2NlcnRp\n" +
            "ZmljYXRlghQvHGgHfN/J7QzPNFAa0q3DwILanjAPBgNVHRMBAf8EBTADAQH/MA0G\n" +
            "CSqGSIb3DQEBCwUAA4IBAQAlbxUF4Eaz0tXSo7oM02Mt3YqhuP7XZpZE5KYpn5qE\n" +
            "utYzJdSJeMsfUpZcmv1pbZ4uepxgBxQKssKRmglzEMX2wxl9WyEPxKkyLTX+XCX9\n" +
            "Vd6IBi5Pft6v2u94bKlGZKigNojGfbzXDYuSU6SAud5GD77RM1vx/pPAa2eG8qSX\n" +
            "OcGQAtHrcSAvl58IqXAmci3akNKN4G5PxNoze5lQ25umQbHTlwvOMwFgPSXseYvm\n" +
            "/f98b+Q6lIdklw6g3XWUCmTkscRM+5mvb+1FKHWU8KiXN7CM+ONXjudO8Ixyyion\n" +
            "pBumFgiA2FQXUpunDCv38dccPb8y/EyhRSQyx+Trusted+\n" +
            "-----END CERTIFICATE-----"

        const val TEST_CERTIFICATE: String = (
            "-----BEGIN CERTIFICATE-----\n" +
                "MIIDCDCCAfCgAwIBAgIULKbJCRQZ251UjMFdqyrwfyWlh60wDQYJKoZIhvcNAQEL\n" +
                "BQAwDzENMAsGA1UEAxMEdGVzdDAeFw0xOTA5MDQxODMwMTFaFw0yMDA5MDMxODMw\n" +
                "MTFaMBsxGTAXBgNVBAMTEHRlc3QtY2VydGlmaWNhdGUwggEiMA0GCSqGSIb3DQEB\n" +
                "AQUAA4IBDwAwggEKAoIBAQC1rTHTp/PvMe0P3Pb76nAr/xrdXehrvgYldPzv7I2W\n" +
                "nR3JHMh5Sk9Xmra27Jd61e5257C3EKa8YpnAISRvPLUUqcFzvncQ6BYHIiLwdRpX\n" +
                "eNYh7OlqFQQ2/6SgeJ2OXWj3QqI9Ih3QVHYAnvMP4cGoVJNpvkotodE+jHfjFO2i\n" +
                "jmyVkKGqNSyy0YKvF0OcNjr2RVvG/jyFmMEzIqSJtJ1wZtEZITZ96ZaLdTD+6GgO\n" +
                "xuCHdl0nUPI+m//3hSon/f9YNO4FTfMjUAjuwh7r5fX+VdEypWcyn/flwfp91xZ6\n" +
                "lLtjJuJ+/ZcUH+Dcdwm1VEp9gVrJLWGDn5hdrN2Cay5HAgMBAAGjUDBOMB0GA1Ud\n" +
                "DgQWBBQ5g2Wq21t5ktgBaHKCXUokgFN5iDAfBgNVHSMEGDAWgBRio//8/deqM5Bw\n" +
                "98SgEqfAy4btJjAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCROXru\n" +
                "J3h8hKRsS4EvLpwyfEPp5A4sdY8mPzX0XGNTfnXnsNPvprDmwDdlTaejyMbK7vD+\n" +
                "V6R2XFr+D5XwkMqa4sWRJzPPj2tydCCb695Ozb2v5vWteiaUbIJEJuO4VU7paKPa\n" +
                "HltRvnI8BAv/8wdIm0xnm6hl7Gs3lGQs6ei1AonFnjvGEhg+/H7104qm903Pkcr6\n" +
                "XvHXp4cSK22OeendqVjpLAEnP3GOq5caGIiU8Q/sDbfVS+hh3eC+86fVpsWxn6rj\n" +
                "FSeZU/sNt18vDgxfPmlHsNOdmphwOneX95fAbE6Bj46b6uJg4wJ6yaRtsXspJtgk\n" +
                "7S/JfXIWVECjyaUE\n" +
                "-----END CERTIFICATE-----"
            )

        const val TEST_PRIVATE_KEY: String = (
            "-----BEGIN RSA PRIVATE KEY-----\n" +
                "MIIEogIBAAKCAQEAta0x06fz7zHtD9z2++pwK/8a3V3oa74uurG53+YRtx6cqsxf\n" +
                "OwkFj2m9vHMUsRTwWrOdCrZVxHXtuurDPwb1Qvuje502AOXTpxlF0CEes7yUHeXP\n" +
                "NmIaah+MKcEIJz3H89rgHOXq5do2Qn6FM+TiqyckSh9YPDBlmTN20VDGyA1nJfNu\n" +
                "NnKbpPibmkmxxZK16mfEIx83PrgWunHCoykujXBcRNy0DCVDC3bL/nfCXAesj3Mj\n" +
                "R2FfgBM/51hVw/3/RnrEtr0XP7tw7Lmy3+z5/dCDiOHoIY/80Gj0kywvHOf3uvjK\n" +
                "gT9pKX/x9nmSrwuaTsEQxc7pJ4+qvSnXtuNXuUpUEOpAnydVljIc5ooIOOrejJSv\n" +
                "53vgg9BkZk21pgV8H0ljplDzrk68VNJ8mzK1RY+7x0paGQkYvVEA6SKfcfFbIMkR\n" +
                "9mQDRhH4fSP/SjSDij1sMV/FmUh6ekUW+bmS/B5N7/92ggPNd9f0VrW3rP8s2J1x\n" +
                "M9WFbsb5bcPFV1QvOAhPsEZGPo6hzJsgqTSgLYV53/yzBJOUGI632vaaEys1Fv3e\n" +
                "ec4dkmJeT86Q6WDI7ouMoKxRJku5xJcmYdsO5PhpjMFkvSjul3Lsn7Lg/gJguXph\n" +
                "ndmrXXeMGwWVGVU4awTB0IBHNjLv8s+3diFFr2UPc+2en8Lo891uksICKezYGhX5\n" +
                "W3tZ8ChJimIN2MT4qqb6K5bzwBl05HMEG/Q1lGkC07GeOSw6VKcB+DCaku6y0MS5\n" +
                "EgJD0tZlv6EBL5fHt6m3fU7Jw22YOIusRaCY7n4OqoO6hHCkWDgxV6MZ5IM3c9W5\n" +
                "/Js6t/q6g7Mf4yJPYf2RE7zFubAo4g4lkc5OWzE/qnp3yxSQKBo9UnrJyRZXVX4g\n" +
                "OXdsL//VT8LrT9i/oFo8a3GzfQGL7QjrtFS/6Q0TJTGDg0jskOrU/OF1ZBZ9X5xO\n" +
                "Cx/RACO/36It19LsyUXWHyEG0ij+783jZ4pKOO4y3zOt2RctQL0lq+vg3Bk+rQBw\n" +
                "J56V9LrliyqSqmyb6y31dCyvsDfySLXkYzorBg4YFEhIC7rfxYyvjZ8dRf9i2LPj\n" +
                "S4ctyztqj0Iy0Q3QlrAsEFAf/alGcEYbJn+l4hkp0H6WOWXO2RdhSNvM9Hd5hhp9\n" +
                "GrgOLfoDlTBi9HCE+X0TGtjSfXUWf5gdJtZFUj8g/19HRU81DqV7o8L/IAW+Xu9Y\n" +
                "s/YxfIXcwp5LJHg3ywx1lYmqzgjJ6YulP1esXNOS38Jvp+OGYgQYzK8xi/fBjYPX\n" +
                "hiQEJkBnSWwKZCZrBkSbctq+EdICkMh+P/3AZdAMZjs5bjcwJlH0OjCGJty8wApJ\n" +
                "h79E8hUatkvVALZmmXOxRv8KuBBAFcnCueG2+SqcJfue6eboMmgmVFKD5yn67nAr\n" +
                "EROFbbQXvHur0vHAyh4G9i/Q2dOBaDGYuqLyao3ZT5JNy8jG56fbzhrRgp380Q7A\n" +
                "LyLCfLCDx4F8h0jaX6wmZ8R8lqqe8atosclzePTAKstolVPftmYR3KOpWqBiZxCx\n" +
                "9rAillBkJoyKsphGKFGfPk7HUGaLYx/PGFmFa8CXqdm3S2x9Zkf=\n" +
                "-----END RSA PRIVATE KEY-----"
            )

        const val TEST_PRIVATE_KEY_PKCS8: String = (
            "-----BEGIN PRIVATE KEY-----\n" +
                "MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIkHoc2YFDSu/EhPl3Q\n" +
                "4zC09GQs/smRhgtHRcyHGdEw9Y1zXJ8NfOu3ql5MJNZ97KJ06W6907x5EBl6HBrI\n" +
                "DsRNibaGPldBtRtm0WWxhoOdZBoLkExk7dZsNzU5/6XGGd0joRi3uYr8UH2ipoYo\n" +
                "aQlK8ItVyuqNweIvygV+tyumRn0tRVIWnKsaSXQu1RQhT9UXRTV9WkjM/fzhRCLM\n" +
                "nJEXml3dgCZGSjp06FjiOXE+5rUCmApxJz1heKl+C//6FAIR/D8Dfg4VdrVivpCg\n" +
                "Bx9s2yk+bdFRUqZbO/BAbWz98Ux2UioZKwj+/nZzb+WycSe7RCC5JIACZiARg2RJ\n" +
                "gC5Ozr9GIPCtMpBifAeXOkELlYjgmy27kyA5kXcdwhgIvykykarjZuerYFtwGDcS\n" +
                "LBT3+wdNH0FbAkc5NqVz+9JoW7Hi5Q2K9u/PitZ4vYORV85GcHZqE2Rg99ubMZId\n" +
                "cWO0CdE59fVL/6Twl4KoZ9Xr+jIlmcaL/JT/qCQfNtyUH3XGUaIph5IqoTTzRlHN\n" +
                "krgfRZw92Nux+jwTMaroEq9hiH9gcbBWkkH8bUpU7c9tyNiqY3z8nvGCLXQrQItn\n" +
                "l88J6q5AG6kQAaKpBJqryHb2Zdz+lVcAvGt1PbRTbDAoOhnHgCsJ3uKtC6g1rEoc\n" +
                "s8yHT+d/Rltj3b6UugZmpNGUr2ALOQkv1WZUTQvZ3WgWnNgxbNgiGwN6R14s8dGD\n" +
                "vf3409aevkt8e41CiCfHSVq7nCIliFuyfcljJWzPBCXewBZUmTZnvJsYWHIvL74w\n" +
                "Bm7TN01LR5agd8s7Qon673A8gNzdjwsNmos/RdZG92OGp4/AtnCs4r2a8q2rZSWx\n" +
                "J3ToZqEz518AINHEjKB1yVbaGcmjwrkfFkbzxkeI//ILSIiVQs231DKYon08NN7H\n" +
                "le8c5P/zTJU+JSnRBorSY5A9IiFNTwcZpcRm/+QJuX/QjfvsFj+P9ipS+XIpXW5x\n" +
                "kM+/vSTg7WgENFMs65YTqaKy4ogq3y5nTXgnmSaIEpoLIWeOLuVmZRjoBOq8az9c\n" +
                "sq+lSkTqUNkAsFGRGvzC72ceE89XPl7236e4N6YSiMVcaSt1hf99zEFTv4Q+iHJH\n" +
                "N9S7d9J2N+n3gePGmqDP49QkBT4fk4CGnyrlHvUgikKetxrOtxw6B8LlGlZswx1L\n" +
                "2hakgdS7EOf/10M+QsTkUib/Sx6GSY3nffi75g7cRBMIYQ508kX+8xhLApTWkzVi\n" +
                "reCWyxrCJpIrhAc55o1kXd6ybqfs1+18SVbGoyynyR2zHgDBC7s1Xvi4Mx7TY73S\n" +
                "ej8o/ntoYsvUdzAcs9t0e9SZ+zHdwGA4CG1wglZGqrWvStX5uuySYDC8/G1HNJAL\n" +
                "GYIP7jK6TaxexhNglt+0jGhdx9uLdHGSn64RyIkCLUVu0LdmguwbGgbVV8vfVSLx\n" +
                "qus5dVhoEiEUiJm7pNlSc01bllXtNvUHmpQcPwRv4AyG0vcDexk70pLri/dCTc/e\n" +
                "yJEkPIAGNGzICGsurBc8mlLraTUY6+BkdGsuKtuXNGZay5qkipEt0zov/zYKwCA+\n" +
                "u+e6cxhJPij8owDqN/xr3e==\n" +
                "-----END PRIVATE KEY-----"
            )

        const val TEST_PRIVATE_KEY_EC: String = "-----BEGIN EC PRIVATE KEY----- fake\n" +
            "MHcCAQEEINZ8J4RQ62sqtAPjcXxDjGbXxw09K+9elaypfMlPwsNWZRJFCjKDDS90\n" +
            "DSWXfsOnweLquTyb1VCymUu4C0a/XX6DTO0s6lISNeNf2eg8W8Zkl9EekcBtXntm\n" +
            "1n8B8C2dEZyQol7pA6fbGApR4WpY10kqBF==\n" +
            "-----END EC PRIVATE KEY-----"

        const val ENCRYPTED_TEST_PRIVATE_KEY_PKCS8: String = (
            "-----BEGIN ENCRYPTED PRIVATE KEY-----\n" +
                "MIIE6TAbBgkqhkiG9w0BBQMwDgQI6XcyFAMjPKsCAggABIIEyJX2fDzS2EchDJ01\n" +
                "UtJ1NN7SFcOXzMtiBGvUjupuuinbiwbi+DEMBT1z/bfwrFBflvK9T3k1iKsV4B3Y\n" +
                "a81HpcqEpowiS0KjVHQzPY21ot1Otjk75JcsnYU3jHJxinU7agmaWcmweSX4uE1H\n" +
                "N7lGNH/BTgx/XndHc+XceQXUHcnw5I87W/YgQQBHjTN3ntqwoNGJAsYsGD6uUoCd\n" +
                "/a4LKFiZjVfs5FgoMT6D5CwteoOmjDKfT6ymhaHtvzCBm3myt3kT1Rr04rC5b2g1\n" +
                "m2whiO61vTJwqrn8PjXVwpY64sdEglWDhAPomI5IJXR+GY0xb6aVt8GIeHCqwsvd\n" +
                "9SPpfWItFHzS/IIFPWUQ4BgzE4jnWyl5bOVHWev2ch+rtlB3s4nItgewtfnYKIm2\n" +
                "vIMBGj9ACB8v7GpMOm94hzwoxZzJMzNO5J/dqY8s0Y+mm5e2ePMOe5mR7vMpBAAl\n" +
                "4DUVnznqfcte41Db+P4mQvyjAe7S5Ir/ScI4ZBcpc7PxOHDgWFeNjonwsAr10s7t\n" +
                "4OTmlTotlXSe64GC72OYkCo7HTD7tU9b4MvmGKK+YfaLeturRufddn9KfETOwYU3\n" +
                "nudnqr5A3P0ahAHDnMnFtd3GX4JuEsg+7OHp1iH4VzCmLSTpJI2PR80vhO4g1Nft\n" +
                "24DnO2XmppE0RkhlvuNQMAaqiOduZbdu+zqRG7qrH6UVtO9yrdYqimyl5s0BncLK\n" +
                "90iPLegUscAEUOg26xxn7dzXRqrkgKhVNcJuFQ9Rlqqgo4/3gG7oydmcZ0s57vXG\n" +
                "hjnJSxbjqkd6rbhuKyh1WIgmUz+h2HtxreqkVHd0gncjoywFzntzvIgb//mwAbei\n" +
                "6xbgSJuuU8PN2mryx/j9oAsxtBJ7auf2lS1lQaBRI7l4fUeQZI77NMzrWiB/spO8\n" +
                "4OFBm0VN+oBj0rRmjj3kRtq0AIF/6cu/RkvLWGufGu1S95Gwgk+FA//2X2T27M8g\n" +
                "OiPZQEE6Ymt9ixfzUHowAlWb2EfG/YBnHBrzypzWiukvYuuQAp3nCZ2uGUTj5thd\n" +
                "6unJrk6Z+ucfg46E2z7zXmTKlDUNM+6zYD1B66Tplr66A6Ra+Jq3sjeLbinDp1lF\n" +
                "4YmkUE62dO1Ns33bKIKU5g9qzVd+jzJsmViFksYm1/qTwXc285U/g+bLFqiiWJYp\n" +
                "E6J6C3a24GrwIsBgmh1RrwlU8FHNOSVJZSA2pdo5YsSzkhTZljGsgGbc0qERd/rf\n" +
                "86S8Y7/w0yiJKUscJrN9Y95kLV3ldQSpOvsECHEwXow4jQQ06ZgkLwNL+gcyXqkZ\n" +
                "IIEUM+TIXFwWq3gf1LgIA9Yjx11/0j7hI0KZWwNTd5xZA1rv53lJYiYtVtBCJ8ll\n" +
                "P86s3C8HNaIMrDmqTqwGpfrSZQdjYpNJ3QwhMuLNyk5td9oXF3DKoBc1CDmALTSB\n" +
                "ec/lfO3xM+/TlcTbi49fw0PUwUcJvlhXko4GkrSGGO3dSd7ZKHzK76ENK0hPaEW9\n" +
                "Jtn9+05pHqZz3RXxAkKc9iAYJdwNNE4++QDN3xlNSz4lN83UwpzoxK4jh4gngZp/\n" +
                "YW49Si94MZKwQPi9ZUClBoEcg5yHBrcGUfng6V4no3aK+4d8O29CuxM5SG3a8Qu8\n" +
                "0gYZxzCMuD8xe/suAw==\n" +
                "-----END ENCRYPTED PRIVATE KEY-----"
            )

        const val OTHER_TEST_CERTIFICATE: String = "-----BEGIN CERTIFICATE----- fake\n" +
            "MIIC5jCCAc4CCQDiRd5QC3sIpjANBgkqhkiG9w0BAQsFADA1MQswCQYDVQQGEwJ1\n" +
            "czELMAkGA1UECAwCbnkxCzAJBgNVBAcMAm55MQwwCgYDVQQDDANwaXYwHhcNMTkw\n" +
            "MzIwMTkzOTAzWhcNMTkwNDE5MTkzOTAzWjA1MQswCQYDVQQGEwJ1czELMAkGA1UE\n" +
            "CAwCbnkxCzAJBgNVBAcMAm55MQwwCgYDVQQDDANwaXYwggEiMA0GCSqGSIb3DQEB\n" +
            "AQUAA4IBDwAwggEKAoIBAQDmY9+PKJCe6Isk9hV4czDBE/iOu0Zi3rvwIK6dWbJq\n" +
            "j2z3lCJgmsLTG16r4kYXsKwjDfviBFhny0KtGfre+vizOeCHNf+RtIJ2W7bVmZoF\n" +
            "HGsv4niygmmYitlLRkVMBZ08SKEx7X6/TjsFFZYqt9GEVjwJEjSzDsb5HOEsFwgS\n" +
            "u9g7v07gQ2J2+AlsqLmM/IdpAJH1dHIha4wnxZaSlIyvKTMgs2RmOrFEEk1tuKqz\n" +
            "P7XdtWZC1DGxVPpNjyCHjVmcK9uI6nQU6S+FWObOngGYyaOw6nICgPCGmLUt7adK\n" +
            "MIZdYifjC0d4tJFIbKtHT9QcKOYu7EeAD6YVg6Xz8VHZAgMBAAEwDQYJKoZIhvcN\n" +
            "AQELBQADggEBAKT7jsM9LB8tvxRErm1vENAiO6rIkqcSO4+/cnaOISOELu8A0mdi\n" +
            "9AaI9hE60BNeJvZdLuyQKOyIC3+ky2WCza+J9ajoVR2GWDDz60j2YIArSjFhm5Bo\n" +
            "YfnFyvT59b/QRy4Uh4xlHXRPOGfg5vC99tk3omgyV9w37sUzKah/BkwLhtqkCVGF\n" +
            "HIkal0qkT084giO8S7pV0MO3ua+cmaSr8C1pRsRX91uFf8alNTjV3eeobeOb8xe7\n" +
            "ZW83Uc7T0H06XgTr4pH4ploR2hm8d36NlUPaYw4UMmCidYbNliuY35Zc8PUApJGV\n" +
            "PTH1yYdWl8KRAujeBsA+t+pXHH3WxZac4nY=\n" +
            "-----END CERTIFICATE-----"

        const val OTHER_TEST_PRIVATE_KEY: String = "-----BEGIN RSA PRIVATE KEY----- fake\n" +
            "MIIEpQIBAAKCAQEA5mPfjyiQnuiLJPYVeHMwwRP4jrtGYt678CCvnMTkfr6A97nP\n" +
            "DMHv4YhyI+Sak5zavx886NeUB9MWZqW26sN5CUEvgsT/XtxAJSl12MTheUJEi+P4\n" +
            "WrVetcAml1IENpadvOEbiU4+I905ZIxPQppEQOA6QNn5sc3r+QikLzRzwSUKU12b\n" +
            "2MkAtHRQqbQ8vdmessxQ1qziVaoqg6hAwumOKYPSKSjMNSWBKbJJeTUiti+30hFR\n" +
            "gHEalKA5MQ9Tn79cJfvXYfp3GXqCyoeYHQ5IfXsueAXakEIBhfX2MO0utGUbJUuF\n" +
            "1esCKrSOXdpli4/tvAXZvlADQN+MbarK4/qe9hXxvWLuiqItZYzsQuAbtNyhaxrW\n" +
            "gfAPUF9OWRoYvr5l/2AT2Wy60Hd7z335X+UVbGcCvj9snEipTj4bkO/aRz0NAxX3\n" +
            "7H1cNXMgkoT40eK/ap+zwM1AMcE/5vDGaA/pWK0N1DKK/YlDrQ1DMyBp2wydjcjm\n" +
            "1YDjapP83iiRTZkQUBLgjoPW3BIdyDzsSJZzmt2ztrjEJFIf+VeYD9yH1Wv+7dEb\n" +
            "L19jHVPkEo8GFsHo22W4cnzy0UvwOT8ykTD/Cn3eA236zA+c6DXp5vDH2etXV7Q1\n" +
            "v8gyGLraUJbBQTosKK2+aXnWsSeHI4qcPLoQzepX6yuGn83jsR5F+sInR+fA+bdc\n" +
            "/w5yNnRnYrkGqx4urYdKjZUdw3N0Lwoxz0UI+P+NL1kMfxrTx+1R7QGJrVWqNHVK\n" +
            "C/XBBOJRLPL07NFLELSdY/8eW/IfRQLBjGCEYdRUs1UWqBOLYja3wlEu6QivEg3O\n" +
            "H9AAjbilzhGKpmVTAfKsl21s2d3Hj+bdMZfSBICXAfA0cysl4f3XMSSGuizPwLdB\n" +
            "jimPuSfxj/SiFvHGnF9Tr45Ld/Pd+raIfilQMsk0ZEqtJOUypcMim2sR9O1vxHu6\n" +
            "ZIHZ6Tt8Bf5IIX3TwdogfFgYIzg0yWCzplWwLUPPraOgFHon8gWgRh/1yrgB98J6\n" +
            "Ni4RGXRKWfv/cJrv5EB29q0kNLFavSDXbUl2r9KoibXaLEMNou/NShd+f3BZGQVQ\n" +
            "drtXJ9CcQ2B5jDO1J6E24dUtqa7BB87pqKwIraaMmE7ck+i8+vhpCzvxBW2G6i07\n" +
            "HnhuYeGT9ApYbJM0Q5WwgWfsgzmehjPEfGm6j7Sx/iCtsZVp2NPbIQVEh22EWXkc\n" +
            "dmI3SkqL/PZwWmgtuyvZfhyBpYsfKQ1ZxPy6q+kBmhP9+CXBE73XCBObENN0ImWo\n" +
            "esCh3fdHoitejgDVA3/zzVsovRbxVvg7xG4fmIxjSzkIVGRjbK8UcWCYM5SdDZlc\n" +
            "21KkojnEIwdjprJoU16c+YhchI/aDxQR0s5rw+BRZCmj1P7O+yc6VP3K8/c8jkJa\n" +
            "4qquTrfiitcdB0TtyL7jFNUs9fO/cEljyMkyMU8b/F78WGJsI7yj0ohnBHKJDzta\n" +
            "LaeA9d8D1w4koacRiA3MV8VnaKn6yXm79MMHoE2RSKXb+B2heqHQPSlCNGaHLgZT\n" +
            "3UNXvK/gEBoaaoFlUAWYWQ91YIMlFb2GFCk6OOU6m4fQ0bX5M7+U624=\n" +
            "-----END RSA PRIVATE KEY-----"

        const val OTHER_TEST_PRIVATE_KEY_PKCS8: String = "-----BEGIN PRIVATE KEY----- fake\n" +
            "MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDNY7+pIZul0vkC\n" +
            "9yx6yNbup/eab1bR4zHQZy9GVnllK8j5smeewUrjv78v3vKVtlXTllFoQTFDp5xB\n" +
            "AUaj+WfeOOpqZy+FuPm5Q3XtcfSkkdbS1UUpTjEmITuNsucKSR27lhQr5G3/qZbS\n" +
            "YtSYi2PEWsrpJkEpgxU7WIEoBlofa8n2L77iN8d3+ZgNnOVr/muXGco9XsFTP8zS\n" +
            "mvjpIidXkaGdR5ZnwGhEFu8ISLkdK5UlWCtd5LCYzMrhyHvAvOJcv3nv5mjZ3E+L\n" +
            "yeaHKptsoQBv9AoqMKnZFyzm3DxwSMsszEUmP8E7CorCPnYzi0hCcgfO3MEcx9VP\n" +
            "c5ah1nxYvkKJoOTIeLJTaRc2wuIGlTVsgu8upg84/Pf/QSH2Clz/nPsPzk4qFQQH\n" +
            "MjEi8WKNs6Zui92nSeJrSF4gV+gaKqAf/KlN9DRmD6ktnzRzLk9Q/4TIz4TBoa/4\n" +
            "rBA6w0IRNxN32L/+I4Bd7XMcKrvMNrkRJl9p67XqlvWYKqT4Lw7RMX31p5ePojLe\n" +
            "V9jZ3johq2yUjSlLenO3aOHsNahaK/10j9qh8puGLY7xeeunZhkvimp60fwomZaX\n" +
            "IFmtO268xbr/jh9nx0pyuvgXQ9PKsNpdwOGgzywZOE7PFg0gbNYRsK3U+JsDoHZM\n" +
            "vMZZf6EnzAcagNd1Tg2C/Vq7TlS/7A54sza17b4vkSbDpkkt/a8yjMzWKkxKUYYp\n" +
            "opLaYif4w3+VeqylEJB09UinkCaSTReb+uSL8jdNAQqcaNqF647zQcmq/kED6jGL\n" +
            "M+FPisCBRXEFpuoYfdEKRNlhKNyv9d+P5go1Vapn+elwn0VLfTY3+ECbIKyu9cqc\n" +
            "8oCuXPhUXGapE/Ozj5nEjpg019TycOha0FHUB+A61XaF064uO1CNksv/zYQO8l12\n" +
            "v84hn3NJ4Fj5VgL1ueDtUK5TQeiSKYeNYIa+FIkb4GwzhrN8AsIymxmqkM480NOi\n" +
            "B5tFKcjlqp6LAJTS2WoaKbs3o/uv5eIUqahXikTChHl0Zb2qI1Mq566ATGplU5kx\n" +
            "xomCy1u1ErGBitgt4Et+h+H+Fr3lEfQ0R783AO4oG2oyfsrk4lr1WLcVX7l1VOvQ\n" +
            "66W9aaweFHrDUUiu8n498kTnf/JTGt8MXIbhoHDH0QkZe6mDRMI/NHowzpkvC70D\n" +
            "gJC4T0IJDn19IwYJrcvGBg9dnhBMdBRra+2sdIqprvy3NbnW2811Jmd2u7rmTYAj\n" +
            "vASrVglB3sjg6aJ3qOEFtB5PvEjAJBcT4Etq6a/o3oeJ39KEzMFtH8D3kPNjgkuJ\n" +
            "xkOQv0pUUeY0cAiutHHrjIVHSxWPSdxUaMm2lg58emE5Dv7pUODmhzb6HRHXQ80s\n" +
            "NFyJrGlJ6b8A6Elqzkq4J/zj9kNt2AA53hunDmgCJeFrUxDkCadgdbXcif63pW2v\n" +
            "RCDVclDcMdN9TNWzqa9/GRuD0PgBjc4daIKJ2syLfSzi6FAQnciNmGH1gn2qjcss\n" +
            "zrqSEviWsZJcIjA4JzGSWal3KTqzQec1cOEEDVeqO53ORIldiAWhZW9YLmMMRB9d\n" +
            "n10Ykvgc/W1N9+GzUQMw83YDMM==\n" +
            "-----END PRIVATE KEY-----"

        const val TEST_INTERMEDIATE_CA: String = (
            "-----BEGIN CERTIFICATE-----\n" +
                "MIIDBzCCAe+gAwIBAgIUQadunJ4an6ikdvfpZ9Fu4qY/zNUwDQYJKoZIhvcNAQEL\n" +
                "BQAwDzENMAsGA1UEAxMEdGVzdDAeFw0xOTA5MDQxODI5NTRaFw0yMDA5MDMxODI5\n" +
                "NTRaMBcxFTATBgNVBAMTDGludGVybWVkaWF0ZTCCASIwDQYJKoZIhvcNAQEBBQAD\n" +
                "ggEPADCCAQoCggEBAJ4vhJfFua/ddXTz6AMXdgmW1HNp9nH7RVLlYum+Zjij+v43\n" +
                "GlW5Exg39g1h6CxzXWOuWiqyKNzKiNrJ+uxkILYCIII9Za/VbqRSwDIViqHoFDXl\n" +
                "cy8Hfw8PL2Sur9qkml+au0+rGbMbseABRlBJR7CC8msfTw1bCvx9zEbmxRmHsa5t\n" +
                "qH/0SFElmB8qrR/ziSaV31N7G/L59Ib2dWm4SSLAnsUdYCInNZgHsyFS24RMlufw\n" +
                "6VnBfvONA4MCUGtupu4lTDrv1TfcvVhbx1Etvlj4ZwggjRrrNkwrknsyr6zySuNF\n" +
                "WN8Hz5MsZB5UgqrRe7iQzQ06o50ppOdgb2INTjkCAwEAAaNTMFEwHQYDVR0OBBYE\n" +
                "FBWaRkbYbTLEac4A28MBiuMWtgyfMB8GA1UdIwQYMBaAFGKj//z916ozkHD3xKAS\n" +
                "p8DLhu0mMA8GA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBAEzQZApl\n" +
                "V+uBmFg+iPEY9MBTIAspiufNOhR2RcBzxZm1ZvZKRgBZ1Ru7EiTowgPsQVdeQaYm\n" +
                "LE80NU3heJ80sIsmwCKIIv+/GuLN1KUXdr2rfxaH9SysG8ECw8tUGqWdJwuPLVDc\n" +
                "xoq6on696eTYfyqyQ2++1LXvA9cvIddnrSPgrZNoB7jTI4sKvFq+DP87Bw7Cw4Ns\n" +
                "q+dLSlroEgN98g8pBiznQYGf7+//QcvpiYCOkr2HTU0ZV9zyEyuANv7u5Bd1H0fO\n" +
                "e/tTTu22bcLEjn6KUuM0OTs2N9k0FoKHJ9Kpw3MgAjJxj3Zzo+sGdDMaf7jyek4x\n" +
                "P1Lt06rls0/FIvg=\n" +
                "-----END CERTIFICATE-----\n"
            )

        const val TEST_INTERMEDIATE_CA_PRIVATE_KEY: String = (
            "-----BEGIN RSA PRIVATE KEY-----\n" +
                "MIIEpAIBAAKCAQEAni+El8W5r911dPPoAxd2CZbUc2n2cftCrwGd6E9qtJB4/wwt\n" +
                "wAfsPiG6RqlmwYzCl89gWczR4mle4qf90tPGUvQoOr5rZ0GUaaxapRgFiAvhQzUw\n" +
                "uAM/BG2DBF8R3OWPO9S6P2wIWVPr6DaJNvarwuicdM7gbQmU/k9uqjWvKAyWYp1l\n" +
                "Z/AxGFjwvFRgv/zzGzdEL1lO6we1xZB6GeairnoZud6Aoii9tpinPkDldiFY1/co\n" +
                "oCd+631ihqRij15r8YkiyQ/Lh1f2KqMDyP6+vTPaxxFZDJa6xhkTAawaxlRt33PE\n" +
                "2BbmrSmbZAcQoei7lRpNsSYWQahG29LfIU8NkXsbRbAayKfZkEjX3GBmPx94PPb5\n" +
                "ChI98PoYr/GqXGd+JeW2EF844OD35R8ANSq1z4bfOJH5r5oJCo0fQUvY9+M7inwH\n" +
                "zebBL2Jul2lAaDHslBccWiAnKazKdYn2h/HtUY2AfSWJiTY9yIh5g9nSXs4cVIMT\n" +
                "uZ5OiHEHIJkIK3++2yHLlqYoh5WepG8DO6kcKzgJq07GdbJlpEJBsNY1TDMLIgGj\n" +
                "B2eQZP4CCNOlceKgKBaag8JWfL+t3jGlbTnIYuBfMFwJrALNR/SitLIUQClFgnXi\n" +
                "RDLzO77Jsl1LGPZCoxcOUMwDG+c3JHv6rps/m33QL7zHFXlv9v+wqu1NHXtNdA9S\n" +
                "qh0K2BwCXUvhXX5Yhq4P5CW4XT7AcRMTgvg5UOTzeJ3QOy+t/KzBltd8i7zzhHiQ\n" +
                "RJCuulbaN9/h2yWUYwvOyzIZ+DdApH9Y4QcF0mPqrVd3p8m9fVwEeDkBRHTLmfNc\n" +
                "VSy0Ag077LkoeZqz5HJCtnaa1M/2sK+BFp3dUy3z1DUP0pADvhZopTqzEEAHfCK+\n" +
                "QlpZhP09clvXA3I5lZXI7mlm8h5LLMY0t4xvFSOAsOc8cEVT7ut6RpIKtk736ve/\n" +
                "0bYSP6HVQnLnv34UexsUwzdsWhWFC2oAWqk9NNqn7tkd/a6DTupe8sxB9tMA9BL7\n" +
                "X91wJTKzzTOodJPbCM7oW+bMNFjfvOQGbBV6o3bEbIccqX3W+82iugwnApWZMWD4\n" +
                "t4LdsenLTBy3atQMe2lcTtPs04rBnJ7KGGD58SP7dBlNKlOihSS1Pwb8UjtCX8fc\n" +
                "KHcPNzGTAzKW4OIX8/zejwrEnuOm0kQpz2Dc/vU0hLx4xqZacfesDY8D7iGCWl1G\n" +
                "cfqHgP9mwU3dxPUZTD7GDFRGqoh7tluLn+32BnY3nFizGVTuDPrMUUkz24P9Uj5p\n" +
                "WKwnU75m1isCBWEd2NsRZgZeCte7oUkGjwBED2qLK7Gbz3XJtjJorhGEXwu3N2XB\n" +
                "sHSm0vFIyMJ8wQBtSUb/cPD+oU+27hhSu/Vb5sqZP8omJIMz3hJ40xLiYr+olpef\n" +
                "hw5sCcwPHYzt9FdMGAVeC3VXlc8gjreF5kWJUPPcec/Bd1Qsv7YRcSCrewP2bwBV\n" +
                "GDZyIOSGdSroLLW2j6hYlfKbx+NFsUAFwioj0H9WtEJfTMrASq1WEiVlUTcYgtiP\n" +
                "uPkGG7Qw2m/qwJfQ5oUtKKzywiUNzcuw+Lovqjv/XesI8pykDwir4Q==\n" +
                "-----END RSA PRIVATE KEY-----\n"
            )

        const val TEST_CA_WITH_DIFFERENT_SKID: String = "-----BEGIN CERTIFICATE-----\n" +
            "MIIERjCCAy6gAwIBAgIUNUJYQYtlm5xiXRw3nyCCF0yDCuwwDQYJKoZIhvcNAQEL\n" +
            "BQAwKTEnMCUGA1UEAxMeaW5mcmEubWFyY3VzLmNvbSBjb3JlIFRydW5rIENBMB4X\n" +
            "DTE5MDQwNTE5NDk1OFoXDTE5MDQxMjE5NTAyOFowLDEqMCgGA1UEAxMhdGVzdCBu\n" +
            "b3QgcmVhbCBzZXJ2aWNlcyBub25wcm9kIENBMIICIjANBgkqhkiG9w0BAQEFAAOC\n" +
            "Ag8AMIICCgKCAgEAwgnArNvkq50OJ7Tt6ykCwdzzMhGaNCWn6jfiVaY21np6J1yJ\n" +
            "Q+9tt3dmRJIiE31K6JM2UwWDUyNEMM10Y8UgYTObomlPpuCWQRHG5OqixojDosAh\n" +
            "UPgqo5EWZdjktnth8uBfw6OWRd7C118e4eaZj0KMSgcKshe4Ed7z28+E2xkk2rzY\n" +
            "oCuXcx1Zee6w8JnZxUjW60G9HKXL+LpeXjUde+LKRskaOlPA1SM1rcWuNy9FP/IL\n" +
            "4fmZAlkOVSfAkvMmZVFZ7m5df+uE1UsaWZOGcJ1MMWBSMyn1AzCFegyHNHozoy/9\n" +
            "V1I3JkqCATFR2WjjYb3TWoJtWshZBrf8SxGEC2XK+K1MI9Z05f+3g2aAq+vpOepd\n" +
            "RL+0LkE1KhTnXmcWRVixOTQFtnFjE+JE0TtZa1Nh1EOM/thkXRLOMZVUJOe7H2Xo\n" +
            "DEyyORYAJkIStRlKnIkZGq+dqo+Li3SWSCb8XFVMopEHU+jHAwu6oI59RIHcCr5a\n" +
            "lGy7dEIy/vUzH5NSG4Uug7iaNrukLpgShVjUyBnRtumpsbOeNaN5TgEfDndX9jaD\n" +
            "gG228Zrajk6O99EGWQEMScIhESYZPZn8MtYFEBZY2QouWovKP65Y2Oog0DHV4Vjb\n" +
            "466TIUJj3hd8ksmfbhzuz0VfVhj5J8eb5mAAUeQ3Ixv3AXAM6+AxgZ9o1+kCAwEA\n" +
            "AaNjMGEwDgYDVR0PAQH/BAQDAgEGMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYE\n" +
            "FM3wyYyxj8mHCXUW2vMOGqIiBeGDMB8GA1UdIwQYMBaAFChgNGIZTLkzfv0y28sd\n" +
            "FtRTk+vTMA0GCSqGSIb3DQEBCwUAA4IBAQCVReWad0ftmFloOde1B3O/6kUYkDk7\n" +
            "LRN/aEo48j67NwJf4vjFuFYiYqVkFUMlkaX41VKfn70fklu6rFSvgpLxeIW+uKVy\n" +
            "cphIcD8VazzL4iMlBLdVbcOTy8DHFxSu/Rn9N9aOAYN9rajnYb/mPFn5r0n/Z/oS\n" +
            "V7XBZwQ31zVw/sdp7Blf+6xOl0KPtL9OJOCWGCAZ27uQdzzHiOCySer8Tc36PP9d\n" +
            "3snlHidzDXLLuPtVF1FmQriP5aPYr8J95im03ssnTX6L/3oK8mUQvDv71Lhl+0IK\n" +
            "1UJ80CY+mpMagyPCzwTw+mFcebORbBdkDbF9plzahR0ui6bK4/fJMUU9\n" +
            "-----END CERTIFICATE-----\n"

        const val TEST_KEY_WITH_DIFFERENT_SKID: String = "-----BEGIN RSA PRIVATE KEY-----\n" +
            "MIIJKgIBAAKCAgEAwgnArNvkq50OJ7Tt6ykCwdzzMhGaNCWn6jfONeb02gT5q9wp\n" +
            "J+1kZ6HMVmZuh98x1Yi3kBxdivjImY80g2sXhESRjxLnjNwwzvZP7vidblJmydgI\n" +
            "cQQXn9ILSdmHIUTa7CbjJ5HbJo2q785E6pMIX1LvlcwBKfl4pT9w43+k9CAH5YMq\n" +
            "nUXPEa4FLs2r4ZVREvMD80w1mJxQ+SdVlaFew+AqcpeChcNJ7Is6pTtOVe7Zt/nU\n" +
            "2xfJjZfIXnxtWEucVhvm2k6lg+dF6OSLbYCZIP4qeeqRvYL8KVyiYcEBnTFPmg/0\n" +
            "s4U5WAIxSTaO8AAbTi7XddZEMZkZVGY1gUtfS4uR+B9Wb2r55D+6r0YRR+MQLhwL\n" +
            "rs+7vSN4msWxippxNoQYsGiIZTZMc+wj2dZko7Mb8uVX/fpvpvdUyuaMRZe0m4bR\n" +
            "bXPFrbEGTEWpbOjcpxmVcd+Sww+nV8aQvje0kBhFkEZOX+exijc8cz97RrQSVW7X\n" +
            "tpV2mfyN/MqoN2VEV2fyo1ZOQIudOzTCxkiTkcGRAAmOAwqNZKE2EoUFjgWM1cAo\n" +
            "VX208VCUwE3H42mPdJXvOVDWWhCtqeI6UNNNIhUn9XYzmYxDF72Z1uAd0TIR8HGY\n" +
            "994FURCD7Kb0HFLuKjkqH9UXKKr9h7NL8AtOjLF8cWK3Bgsx8+svmt2n2+ohLVVG\n" +
            "yjtPGAujsHGLQb/HNaU6Lg2E721FmiirykuBJY3+Rp1CeVQvWkGGYuJHIEOG72Fl\n" +
            "ynFvkGQSWa6pXRV6DLmgUu2120yIdaxB29fKI13OsRJtnP5/FKVERaB4V2dANXdl\n" +
            "/BNZDvFR65SfxZbIlJZbexQZNUCfLPKENdwnhOFpN9DWjKEcQEzD5TQLXAWvlyNS\n" +
            "klRdebacJZlqwM5a4XtkDXJ97k2hv/f6cn5nnQv/2i++wV838RBgxHSrMeA3MBX+\n" +
            "T5OexHmpuGSxfYI01AezE+cqCG8UkbM+3h71vEgw090Zy/t6U+Q+q0RScs+x3Czq\n" +
            "kbSxZ5EhaFRTtfuXE5x92x1nOwzQmuCOgUBeD5XAg/T+SkZawqTdzC8kS57ibNBs\n" +
            "mlYgbazIct/2sdQH6FZ2j6RbdB0BtKYC0a5YTW3XZO15eTm0GMOKqRDVSKHXt7l3\n" +
            "m7s3hyTAH02vWP2Ds/mmnK3ZftzwLaTgm1n2csO5U9tvXCP24Yvbvx1b11Fx3KTf\n" +
            "w+aY0Uk2/GFEU3y4dVMfF9LmTaETI76uV1IRk7jb7nc3YQA3L+6VKfS6ukH9LT7f\n" +
            "x7qNMKu+Su5cxbkPlWpI5h5uf9CmCCUiurunTxEJ2iVrg/4lMAMf6n1SGXUdIXUx\n" +
            "V5I80SroqVRWh+VHc8SDwwLfNtj+N9HiijMZeECVALJqW7Fm9nkxuoiLNLEc07X3\n" +
            "y+MJhzuc4CY46BszkOixoUvxeh89sfkK4udqmtcOyjeVARU0OY/5k9jdWgxGmXJL\n" +
            "MeGOsldDLCZ0mbcOQFPwvGCDibczRQ1HowoL9tAK5rF8Q8TDwueNul27wiKZvjJU\n" +
            "qe7fj30RbbMYlR7DDzt4WNb+ULqnsafgJGLg1nTZwuWd7PAuAJ2eXqVLb7iln0JH\n" +
            "fux91Nu9Vd95zNjYL8+o9Y7xqPApy0MGhAHaMYeso4Z4CcJe4l1wIoR//DyZ15J4\n" +
            "I8DmnbPSYiFCib4BBlCqxuy2prSNHUYfP4rMM9z2mq2v6mLGxhWPRCm+J5yn623a\n" +
            "raWlIqcsYZS5BtBCaMmkbTdXG4mwgVgjHcK1gy8jH/t7Yr174UgU/UBkkXwHGr00\n" +
            "iUilj0BVYt4W+7yT56QSFWaP8CNe7nsS5OJ3zBfSEmPoJMbvO5B6Nu0iijVuPXQn\n" +
            "ns+5iKkDHsrZnVh0MnJV9Eritmk38Mmr9fqif3Txj36YNIyxwfYo0yhSmIexy5ty\n" +
            "fgIn0mNrLzOwezLxqqJ4aMEQ3J90XG6pZjlaK81wfapo6WhsedtLI22a7Gb1T5SL\n" +
            "PT2YXH4VLTmal2kzBriJXHBQdfd7J6RIdaFkQ5VUsNPnwLCvH48BTvw57r8IXXC3\n" +
            "/xb+vDEzSw3IKAZ8uMqUqzpbvGHZ1s2epXfBbqIO14FhyrShEAW8zOj9N0UaRmHY\n" +
            "YaJ/bQ0NV8/5Dnj7sl/9hA25WSSGpb8RNI+v0hEHpYjDermBVh3/tri56FP6+lak\n" +
            "5BQWCfXL65cyPSh0ITbeK0pPIvF/dOl8UiIYsi+lFbrBoAhZDEtiLanvlklidXWv\n" +
            "JO7zbvCWRofWd5UTXzQ1N+DkJ202e3B33tcwlWsrBvZk5DI6JFnuQ3LOpD74+n8Q\n" +
            "53DkfltxD7z3tqhlLpi+wCFQVbtkFEyK7ZuUVXRdQ+6eUBL8hM7odXR2u9Tn2I8b\n" +
            "IxnsBlsfje3gLm+O0IeP7CenEyiUCcmoGcbLYdow91fqYx6jl0CVqHe7xR2OmDtU\n" +
            "qieUpnzoGzJdtNBM8Z7AOA/i1JKjy/3fBF0/lPGl4bsyfV0H/1DSRCHNnuk/lqYP\n" +
            "CLJxOokujLSrbExRUcMATBLIFbZKnJ7CpTi+VSWQd0e01Nekde4zrXcgbdDMJ5KW\n" +
            "Zvmx+V9auDn/nOMqXUriXDrG942owVuKGsIo4ddHqQjH6FjiVA6YHS1vIIC1ULD9\n" +
            "VMQ2EF5chixYJFhEaJhRsbkxRHBZUQjjUHUqV1I+YsRqpp4eQxswhh3BxaKCsMFI\n" +
            "SbdPRcXYcP0uQKAmj26weEouKG56QwoUTqI13wB0eYarqJ3tFNmbtyMNygmzU8IV\n" +
            "HRPYqSgoXmoI06IQO61fkJ6c8aZWDKdAYg2SKrcUjIsrbkkLj+y2S+1MWGpHhqjt\n" +
            "wSNaTc5Yj1dInD2z3IH3lryCaLYvCtdXCSD2D7MHxICl9ud0sup0iIJZj3PfQIhP\n" +
            "VJOohMblXcsXXeR8NzpYm/LIALoTDVCSuS1FuyaOyK270njaa5gldnsFaQaGa8VR\n" +
            "ra2R8DK1Eg3C4XuGtOkqsB/DC9G0FLYZEi2Y8BRiTsj3sgnA49AsmycWqCrfLOO1\n" +
            "n8Vai+qvQD3RFqJYraV49YyyYE4x+Cwsq0dRNBcT52Ux2Jd7d8WX48K+RgBV2JP3\n" +
            "E46sfeKbyIjTWyUKMgdF4rSWdIH1T86g+euoHbo+YWwN19JUGOilXvlsp+ApkF==\n" +
            "-----END RSA PRIVATE KEY-----\n"
    }
}
