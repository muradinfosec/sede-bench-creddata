stdout_logger = Logger.new STDOUT

password = "690108f94398f07cd426578e7478458d" # $ Source[rb/clear-text-logging-sensitive-data]

# BAD: password logged as plaintext
stdout_logger.info password # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: password logged as plaintext
stdout_logger.debug password # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: password logged as plaintext
stdout_logger.error password # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: password logged as plaintext
stdout_logger.fatal password # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: password logged as plaintext
stdout_logger.unknown password # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: password logged as plaintext
stdout_logger.warn password # $ Alert[rb/clear-text-logging-sensitive-data]

# BAD: password logged as plaintext
stdout_logger.add Logger::WARN, password # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: password logged as plaintext
stdout_logger.add Logger::WARN, "message", password # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: password logged as plaintext
stdout_logger.log Logger::WARN, password # $ Alert[rb/clear-text-logging-sensitive-data]

# BAD: password logged as plaintext
stdout_logger << "pw: #{password}" # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: sensitive data in the progname will taint subsequent logging calls
stdout_logger.progname = password # $ Alert[rb/clear-text-logging-sensitive-data]

hsh1 = { password: "dcb3150f96a7e659560b9c89ee2a74d5" } # $ Source[rb/clear-text-logging-sensitive-data]
hsh2 = {}
# GOOD: no backwards flow
stdout_logger.info hsh2[:password]
hsh2[:password] = "bcead073a2038d35363a18ef6941b285" # $ Source[rb/clear-text-logging-sensitive-data]
hsh3 = hsh2

# BAD: password logged as plaintext
stdout_logger.info hsh1[:password] # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: password logged as plaintext
stdout_logger.info hsh2[:password] # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: password logged as plaintext
stdout_logger.info hsh3[:password] # $ Alert[rb/clear-text-logging-sensitive-data]
# GOOD: not a password
stdout_logger.info hsh1[:foo]

password_masked_sub = "fc817212b3c863230aa5b03fe6bb6021"
password_masked_sub_ex = "dc401937b0f705242ac1b81fe1bd6862"
password_masked_gsub = "d0e7704e83361a5a2b1982247105515a"
password_masked_gsub_ex = "f3f0181f02032d5e1f6450683950613e"
password_masked_sub = password_masked_sub.sub(/.+/, "[password]")
password_masked_sub_ex.sub!(/.+/, "[password]")
password_masked_gsub = password_masked_gsub.gsub(/./, "*")
password_masked_gsub_ex.gsub!(/./, "*")

# GOOD: password is effectively masked before logging
stdout_logger.info password_masked_sub
# GOOD: password is effectively masked before logging
stdout_logger.info password_masked_gsub
# GOOD: password is effectively masked before logging
stdout_logger.info password_masked_sub_ex
# GOOD: password is effectively masked before logging
stdout_logger.info password_masked_gsub_ex

password_masked_ineffective_sub = "bb161999b6b791126fa5a09dd6fa1829" # $ Source[rb/clear-text-logging-sensitive-data]
password_masked_ineffective_sub_ex = "aa298731d5e867881de4a14cb8dc9096" # $ Source[rb/clear-text-logging-sensitive-data]
password_masked_ineffective_gsub = "e8b2078c09364c1a0e4164037097126b" # $ Source[rb/clear-text-logging-sensitive-data]
password_masked_ineffective_gsub_ex = "a4f4291b05545a1a1d3642974056944c" # $ Source[rb/clear-text-logging-sensitive-data]
password_masked_ineffective_sub = password_masked_ineffective_sub.sub(/./, "[password]") # $ Source[rb/clear-text-logging-sensitive-data]
password_masked_ineffective_sub_ex.sub!(/./, "[password]")
password_masked_ineffective_gsub = password_masked_ineffective_gsub.gsub(/[A-Z]/, "*") # $ Source[rb/clear-text-logging-sensitive-data]
password_masked_ineffective_gsub_ex.gsub!(/[A-Z]/, "*")

# BAD: password masked ineffectively
stdout_logger.info password_masked_ineffective_sub # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: password masked ineffectively
stdout_logger.info password_masked_ineffective_gsub # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: password masked ineffectively
stdout_logger.info password_masked_ineffective_sub_ex # $ Alert[rb/clear-text-logging-sensitive-data]
# BAD: password masked ineffectively
stdout_logger.info password_masked_ineffective_gsub_ex # $ Alert[rb/clear-text-logging-sensitive-data]

def foo(password, logger)
  # BAD: password logged as plaintext
  logger.info password # $ Alert[rb/clear-text-logging-sensitive-data]
end

password_arg = "93c1270af4f3e7c43e0bf4ffcb236536" # $ Source[rb/clear-text-logging-sensitive-data]
foo(password_arg, stdout_logger)
foo("65f2950df2f0e2c38d7ba2ccca767292", stdout_logger)

def redact(password)
  "***"
end

password_r1 = redact("35c4003cc6f9a4d33e3be0dcdd002813")
password_r2 = password_r1
# GOOD: password_r2 has been redacted
stdout_logger.info password_r2
