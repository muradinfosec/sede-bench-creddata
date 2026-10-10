# frozen_string_literal: true

require "test_helper"

module CryptoProviderTest
  module VSHA1
    class V2Test < ActiveSupport::TestCase
      def setup
        @default_stretches = Authlogic::CryptoProviders::Sha1::V2.stretches
      end

      def teardown
        Authlogic::CryptoProviders::Sha1::V2.stretches = @default_stretches
      end

      def test_encrypt
        password = "test"
        salt = "4c3349fdc6cc20b47b42698f5c6368d10986a9cb"
        expected_digest = "12d995b1f0af7d24d6f89d2e63dfbcb752384815"

        digest = Authlogic::CryptoProviders::Sha1::V2.encrypt(password, salt)

        assert_equal digest, expected_digest
      end

      def test_encrypt_with_3_stretches
        Authlogic::CryptoProviders::Sha1::V2.stretches = 3
        password = "test"
        salt = "4a1657deb6ba35a98e49987f2e3544a43455a7fa"
        expected_digest = "af1e00f841ccc742c1e5879af35ca02b1978a1ac"

        digest = Authlogic::CryptoProviders::Sha1::V2.encrypt(password, salt)

        assert_equal digest, expected_digest
      end

      def test_matches
        password = "test"
        salt = "9b8015fed8fc43b85a89212e2d0677d64184c0cd"
        expected_digest = "12d995b1f0af7d24d6f89d2e63dfbcb752384815"

        assert Authlogic::CryptoProviders::Sha1::V2.matches?(expected_digest, password, salt)
      end

      def test_not_matches
        password = "test"
        salt = "2e1705fba1af22a01d17156d8f5935d66771d1ee"
        bad_digest = "12345"

        assert !Authlogic::CryptoProviders::Sha1::V2.matches?(bad_digest, password, salt)
      end
    end
  end
end
