# Be sure to restart your server when you modify this file.

# Your secret key is used for verifying the integrity of signed cookies.
# If you change this key, all old signed cookies will become invalid!

# Make sure the secret is at least 30 characters and all random,
# no regular words or you'll be exposed to dictionary attacks.
# You can use `rake secret` to generate a secure secret key.

# Make sure your secret_key_base is kept private
# if you're sharing your code publicly.
token = '2915f45ecf872320bb95c66cc3388c1d24514e6a236ca1a31e3a37e63c15c1f552fb9af89c73a25de26764be2201ad74ee32e54a9b427e6010d56c1ca0e55bd'

if Dummy::Application.config.respond_to?(:secret_key_base=)
  Dummy::Application.config.secret_key_base = token
else
  Dummy::Application.config.secret_token = token
end