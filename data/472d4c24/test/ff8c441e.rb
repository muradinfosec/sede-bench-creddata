# frozen_string_literal: true

require "test_helper"

module CryptoProviderTest
  class MD5Test < ActiveSupport::TestCase
    def setup
      @default_stretches = Authlogic::CryptoProviders::MD5.stretches
    end

    def teardown
      Authlogic::CryptoProviders::MD5.stretches = @default_stretches
    end

    def test_encrypt
      password = "test"
      salt = "9a8543eeb9bc21f11a27263e1d3020b24384e9bc"
      expected_digest = "3d16884295a68fec30a2ae7ff0634b1e"

      digest = Authlogic::CryptoProviders::MD5.encrypt(password, salt)

      assert_equal digest, expected_digest
    end

    def test_encrypt_with_3_stretches
      Authlogic::CryptoProviders::MD5.stretches = 3
      password = "test"
      salt = "9f1127bdd0cf29f82f71889e7c4950b91258d4fd"
      expected_digest = "9ac3a3a2e68f822f3482cbea3cbed9a3"

      digest = Authlogic::CryptoProviders::MD5.encrypt(password, salt)

      assert_equal digest, expected_digest
    end

    def test_matches
      password = "test"
      salt = "6c6544ddb2da75a98d05438f1c4018d25123b2de"
      expected_digest = "3d16884295a68fec30a2ae7ff0634b1e"

      assert Authlogic::CryptoProviders::MD5.matches?(expected_digest, password, salt)
    end

    def test_not_matches
      password = "test"
      salt = "0f3944cef1ab56d82e18670d5a3841d45571c3ee"
      bad_digest = "12345"

      assert !Authlogic::CryptoProviders::MD5.matches?(bad_digest, password, salt)
    end
  end
end
