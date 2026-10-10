# frozen_string_literal: true

require "test_helper"

module CryptoProviderTest
  module SHA256
    class V2Test < ActiveSupport::TestCase
      def setup
        @default_stretches = Authlogic::CryptoProviders::Sha256::V2.stretches
      end

      def teardown
        Authlogic::CryptoProviders::Sha256::V2.stretches = @default_stretches
      end

      def test_encrypt
        password = "test"
        salt = "8e9624eaa6bf45a28e98199e4e0540c35889b8db"
        expected_digest = "7f42a368b64a3c284c87b3ed3145b0c89f6bc49de931ca083e9c56a5c6b98e22"

        digest = Authlogic::CryptoProviders::Sha256::V2.encrypt(password, salt)

        assert_equal digest, expected_digest
      end

      def test_encrypt_with_3_stretches
        Authlogic::CryptoProviders::Sha256::V2.stretches = 3
        password = "test"
        salt = "5e6569fcb3ef12d68d42508a2c9854b04065b2be"
        expected_digest = "1560ebc3b08d86828a7e9267379f7dbb847b6cc255135fc13210a4155a58b981"

        digest = Authlogic::CryptoProviders::Sha256::V2.encrypt(password, salt)

        assert_equal digest, expected_digest
      end

      def test_matches
        password = "test"
        salt = "7e9292efd2ab64c05c03685e6f8824a27722a6ba"
        expected_digest = "7f42a368b64a3c284c87b3ed3145b0c89f6bc49de931ca083e9c56a5c6b98e22"

        assert Authlogic::CryptoProviders::Sha256::V2.matches?(expected_digest, password, salt)
      end

      def test_not_matches
        password = "test"
        salt = "1d9595cea7cc79d37a90741f2a6938a81396e0ed"
        bad_digest = "12345"

        assert !Authlogic::CryptoProviders::Sha256::V2.matches?(bad_digest, password, salt)
      end
    end
  end
end
