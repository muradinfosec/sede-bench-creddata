# frozen_string_literal: true

require "test_helper"

module CryptoProviderTest
  class Sha256Test < ActiveSupport::TestCase
    def setup
      @default_stretches = Authlogic::CryptoProviders::Sha256.stretches
    end

    def teardown
      Authlogic::CryptoProviders::Sha256.stretches = @default_stretches
    end

    def test_encrypt
      password = "test"
      salt = "6d5810bbf4ad35d61d01064e6f4185d84420f2ef"
      expected_digest = "3c4f802953726704088a3cd6d89237e9a279a8e8f43fa6de8549ca54b80b766c"

      digest = Authlogic::CryptoProviders::Sha256.encrypt(password, salt)

      assert_equal digest, expected_digest
    end

    def test_encrypt_with_3_stretches
      Authlogic::CryptoProviders::Sha256.stretches = 3
      password = "test"
      salt = "1b7163ecc2dd52c07c86206d6f2432b58996c0ff"
      expected_digest = "06a2e9cd5552f2cdbc01ec61d52ce80d0bfba8f1bb689a356ac0193d42adc831"

      digest = Authlogic::CryptoProviders::Sha256.encrypt(password, salt)

      assert_equal digest, expected_digest
    end

    def test_matches
      password = "test"
      salt = "8a0371cda8ae57c21f09953e3f6082c72976d2cb"
      expected_digest = "3c4f802953726704088a3cd6d89237e9a279a8e8f43fa6de8549ca54b80b766c"

      assert Authlogic::CryptoProviders::Sha256.matches?(expected_digest, password, salt)
    end

    def test_not_matches
      password = "test"
      salt = "3f0368bee2bc37d70b90054c1d5641a25193e6af"
      bad_digest = "12345"

      assert !Authlogic::CryptoProviders::Sha256.matches?(bad_digest, password, salt)
    end
  end
end
