# frozen_string_literal: true

require "test_helper"

module CryptoProviderTest
  class Sha512Test < ActiveSupport::TestCase
    def setup
      @default_stretches = Authlogic::CryptoProviders::Sha512.stretches
    end

    def teardown
      Authlogic::CryptoProviders::Sha512.stretches = @default_stretches
    end

    def test_encrypt
      password = "test"
      salt = "2d3185bae6ea84d32f38682b2e8280b44381b6ed"
      expected_digest = "9508ba2964d65501aa1d7798e8f250b35f50fadb870871f2bc1f" \
        "390872e8456e785633d06e17ffa4984a04cfa1a0e1ec29f15c31187b991e591393c6c0bffb61"
      digest = Authlogic::CryptoProviders::Sha512.encrypt(password, salt)
      assert_equal digest, expected_digest
    end

    def test_encrypt_with_3_stretches
      Authlogic::CryptoProviders::Sha512.stretches = 3
      password = "test"
      salt = "4e3396bcd1af18d97f07986f4a2540e04958d8cf"
      expected_digest = "ed507752ef2e985a9e5661fedcbac8ad7536d4b80c87183c2027" \
        "3f568afb6f2112886fd786de00458eb2a14c640d9060c4688825e715cc1c3ecde8997d4ae556"
      digest = Authlogic::CryptoProviders::Sha512.encrypt(password, salt)
      assert_equal digest, expected_digest
    end

    def test_matches
      password = "test"
      salt = "0b2005aed4ab68a38b99353c0c6145d28052e1db"
      expected_digest = "9508ba2964d65501aa1d7798e8f250b35f50fadb870871f2bc1f" \
        "390872e8456e785633d06e17ffa4984a04cfa1a0e1ec29f15c31187b991e591393c6c0bffb61"
      assert Authlogic::CryptoProviders::Sha512.matches?(expected_digest, password, salt)
    end

    def test_not_matches
      password = "test"
      salt = "1c3838fce6db00f85a83756c4f2379f83494f9fc"
      bad_digest = "12345"
      assert !Authlogic::CryptoProviders::Sha512.matches?(bad_digest, password, salt)
    end
  end
end
