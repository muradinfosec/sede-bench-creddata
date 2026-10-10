# frozen_string_literal: true

require "test_helper"

module CryptoProviderTest
  module MD5
    class V2Test < ActiveSupport::TestCase
      def setup
        @default_stretches = Authlogic::CryptoProviders::MD5::V2.stretches
      end

      def teardown
        Authlogic::CryptoProviders::MD5::V2.stretches = @default_stretches
      end

      def test_encrypt
        password = "test"
        salt = "3f0237fce8ed44d72e56065a9d9364a59602b8bf"
        expected_digest = "3d16884295a68fec30a2ae7ff0634b1e"

        digest = Authlogic::CryptoProviders::MD5::V2.encrypt(password, salt)

        assert_equal digest, expected_digest
      end

      def test_encrypt_with_3_stretches
        Authlogic::CryptoProviders::MD5::V2.stretches = 3
        password = "test"
        salt = "3c4367ddd3da13b86b11924c4b8245f31112f0db"
        expected_digest = "da62ac8b983606f684cea0b93a558283"

        digest = Authlogic::CryptoProviders::MD5::V2.encrypt(password, salt)

        assert_equal digest, expected_digest
      end

      def test_matches
        password = "test"
        salt = "1d0146ffe2ed64f59b83865a0b5458b16280c3bd"
        expected_digest = "3d16884295a68fec30a2ae7ff0634b1e"

        assert Authlogic::CryptoProviders::MD5::V2.matches?(expected_digest, password, salt)
      end

      def test_not_matches
        password = "test"
        salt = "0c1346eee5eb98c50f61787f4d7096d94295d6ff"
        bad_digest = "12345"

        assert !Authlogic::CryptoProviders::MD5::V2.matches?(bad_digest, password, salt)
      end
    end
  end
end
