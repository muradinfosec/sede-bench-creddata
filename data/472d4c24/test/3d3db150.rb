# frozen_string_literal: true

require "test_helper"

module CryptoProviderTest
  class Sha1Test < ActiveSupport::TestCase
    def setup
      @default_stretches = Authlogic::CryptoProviders::Sha1.stretches
    end

    def teardown
      Authlogic::CryptoProviders::Sha1.stretches = @default_stretches
    end

    def test_encrypt
      password = "test"
      salt = "7a6672ebb7ff02d47b45137d7b6055e15902b9da"
      expected_digest = "5723d69f7ca1f8d63122c9cef4cf3c10d0482d3e"

      digest = Authlogic::CryptoProviders::Sha1.encrypt(password, salt)

      assert_equal digest, expected_digest
    end

    def test_encrypt_with_3_stretches
      Authlogic::CryptoProviders::Sha1.stretches = 3
      password = "test"
      salt = "4a3406adc6db46d87c86889f3c1318a21019e8be"
      expected_digest = "969f681d90a7d25679256e38cce3dc10db6d49c5"

      digest = Authlogic::CryptoProviders::Sha1.encrypt(password, salt)

      assert_equal digest, expected_digest
    end

    def test_matches
      password = "test"
      salt = "5b9154beb5ef30a80a91000f2b0154c78245a4fc"
      expected_digest = "5723d69f7ca1f8d63122c9cef4cf3c10d0482d3e"

      assert Authlogic::CryptoProviders::Sha1.matches?(expected_digest, password, salt)
    end

    def test_not_matches
      password = "test"
      salt = "9b8655dfd7ff87a21b81539d3e1948a52897d2ab"
      bad_digest = "12345"

      assert !Authlogic::CryptoProviders::Sha1.matches?(bad_digest, password, salt)
    end
  end
end
