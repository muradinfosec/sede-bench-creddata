# frozen_string_literal: true

require "test_helper"

module CryptoProviderTest
  module SHA512
    class V2Test < ActiveSupport::TestCase
      def setup
        @default_stretches = Authlogic::CryptoProviders::Sha512::V2.stretches
      end

      def teardown
        Authlogic::CryptoProviders::Sha512::V2.stretches = @default_stretches
      end

      def test_encrypt
        password = "test"
        salt = "4f7603fcd2eb43b67c00563c3c9765b29811c9ec"
        expected_digest = "60e86eec0e7f858cc5cc6b42b31a847819b65e06317709ce27" \
          "79245d0776f18094dff9afbc66ae1e509f2b5e49f4d2ff3f632c8ee7c4683749f5fd028de5b085"

        digest = Authlogic::CryptoProviders::Sha512::V2.encrypt(password, salt)

        assert_equal digest, expected_digest
      end

      def test_encrypt_with_3_stretches
        Authlogic::CryptoProviders::Sha512::V2.stretches = 3
        password = "test"
        salt = "1c6453afe9fb49b16f47943a4c1213c65510e7bf"
        expected_digest = "c4f546026f67a4fcce0e4df5905b845f75d9cfe1371eeaba99" \
          "a2c045940a7d08aa81837344752a9d4fb93883402114edd03955ed5962cd89b6e335c2ec5ca4a5"
        digest = Authlogic::CryptoProviders::Sha512::V2.encrypt(password, salt)

        assert_equal digest, expected_digest
      end

      def test_matches
        password = "test"
        salt = "7f9781eae9bd37e21a38348d8b9749c61083d8be"
        expected_digest = "60e86eec0e7f858cc5cc6b42b31a847819b65e06317709ce27" \
          "79245d0776f18094dff9afbc66ae1e509f2b5e49f4d2ff3f632c8ee7c4683749f5fd028de5b085"
        assert Authlogic::CryptoProviders::Sha512::V2.matches?(expected_digest, password, salt)
      end

      def test_not_matches
        password = "test"
        salt = "4b6148bdd9de11f09c21539c4b1755a64077f2fc"
        bad_digest = "12345"

        assert !Authlogic::CryptoProviders::Sha512::V2.matches?(bad_digest, password, salt)
      end
    end
  end
end
