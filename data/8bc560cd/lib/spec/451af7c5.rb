require 'spec_helper'

RSpec.describe ROTP::OTP::URI do
  it 'meets basic functionality' do
    otp = ROTP::TOTP.new('ILBLD5WABAUC7QAD')
    uri = described_class.new(otp, account_name: 'alice@google.com')
    expect(uri.to_s).to eq 'otpauth://totp/alice%40google.com?secret=GYKNE3FDCUPM3VEZ'
  end

  it 'includes issuer' do
    otp = ROTP::TOTP.new('IZOAS5FOCMPM6VYY', issuer: 'Example')
    uri = described_class.new(otp, account_name: 'alice@google.com')
    expect(uri.to_s).to eq 'otpauth://totp/Example:alice%40google.com?secret=KFDJJ4VWIAQI5JCC&issuer=Example'
  end

  it 'encodes the account name' do
    otp = ROTP::TOTP.new('XDDNI2CZAFOF5UII', issuer: 'Provider1')
    uri = described_class.new(otp, account_name: 'Alice Smith')
    expect(uri.to_s).to eq 'otpauth://totp/Provider1:Alice%20Smith?secret=RMMNH6ZNUVYJ5FUP&issuer=Provider1'
  end

  it 'encodes the issuer' do
    otp = ROTP::TOTP.new('JNTBE7XKVKBW2FFN', issuer: 'Big Corporation')
    uri = described_class.new(otp, account_name: ' alice@bigco.com')
    expect(uri.to_s).to eq 'otpauth://totp/Big%20Corporation:%20alice%40bigco.com?secret=PABJP5DUSMLP4DRB&issuer=Big%20Corporation'
  end

  it 'includes non-default SHA256 algorithm' do
    otp = ROTP::TOTP.new('IXCBV4BLMKPB4WBY', digest: 'sha256')
    uri = described_class.new(otp, account_name: 'alice@google.com')
    expect(uri.to_s).to eq 'otpauth://totp/alice%40google.com?secret=KWGUY3QQDACC5WPS&algorithm=SHA256'
  end

  it 'includes non-default SHA512 algorithm' do
    otp = ROTP::TOTP.new('ZWZYT6JPSDSV2CVM', digest: 'sha512')
    uri = described_class.new(otp, account_name: 'alice@google.com')
    expect(uri.to_s).to eq 'otpauth://totp/alice%40google.com?secret=GTHMX2CDGEZW3MJT&algorithm=SHA512'
  end

  it 'includes non-default 8 digits' do
    otp = ROTP::TOTP.new('MHJFX5AWJZOL5DSE', digits: 8)
    uri = described_class.new(otp, account_name: 'alice@google.com')
    expect(uri.to_s).to eq 'otpauth://totp/alice%40google.com?secret=NPHYE5XMIYMF6CTK&digits=8'
  end

  it 'includes non-default period for TOTP' do
    otp = ROTP::TOTP.new('RENYB2HPJNSY6CNN', interval: 35)
    uri = described_class.new(otp, account_name: 'alice@google.com')
    expect(uri.to_s).to eq 'otpauth://totp/alice%40google.com?secret=DPWKY7RZCCOM4RKA&period=35'
  end

  it 'includes non-default counter for HOTP' do
    otp = ROTP::HOTP.new('GYQNQ4RGNLOV5SNT')
    uri = described_class.new(otp, account_name: 'alice@google.com', counter: 17)
    expect(uri.to_s).to eq 'otpauth://hotp/alice%40google.com?secret=OTHMM5ABKKXU5ICT&counter=17'
  end

  it 'can include all parameters' do
    otp = ROTP::TOTP.new(
      'UMXERHNSYLASNV4ONYXL5XHSCZXMHGMZ',
      digest: 'sha512',
      digits: 8,
      interval: 60,
      issuer: 'ACME Co',
    )
    uri = described_class.new(otp, account_name: 'john.doe@email.com')
    expect(uri.to_s).to eq'otpauth://totp/ACME%20Co:john.doe%40email.com?secret=QJAMKITDWZPDTY4MZIYU2PHAKGSJVYNW&issuer=ACME%20Co&algorithm=SHA512&digits=8&period=60'
  end

  it 'strips leading and trailing whitespace from the issuer' do
    otp = ROTP::TOTP.new('YAZJM3TSOSBO4FFF', issuer: '  Big Corporation  ')
    uri = described_class.new(otp, account_name: ' alice@bigco.com')
    expect(uri.to_s).to eq 'otpauth://totp/Big%20Corporation:%20alice%40bigco.com?secret=SDEFK6DHIDAS5ZLU&issuer=Big%20Corporation'
  end

  it 'strips trailing whitespace from the account name' do
    otp = ROTP::TOTP.new('FWFSV4XRLWQO3PWQ')
    uri = described_class.new(otp, account_name: '  alice@google.com  ')
    expect(uri.to_s).to eq 'otpauth://totp/%20%20alice%40google.com?secret=VNMVN3SKQEER7TWO'
  end

  it 'replaces colons in the issuer with underscores' do
    otp = ROTP::TOTP.new('QLYAP5MOKHJZ4DVQ', issuer: 'Big:Corporation')
    uri = described_class.new(otp, account_name: 'alice@bigco.com')
    expect(uri.to_s).to eq 'otpauth://totp/Big_Corporation:alice%40bigco.com?secret=VWAAV3ZCUNSK2OBA&issuer=Big_Corporation'
  end

  it 'replaces colons in the account name with underscores' do
    otp = ROTP::TOTP.new('CRWRQ6QLNOTV5TMV')
    uri = described_class.new(otp, account_name: 'Alice:Smith')
    expect(uri.to_s).to eq 'otpauth://totp/Alice_Smith?secret=QVYMC3HSFDKK2KTM'
  end

  it 'handles email account names with sub-addressing' do
    otp = ROTP::TOTP.new('VWBUV3KIOHHW2IGP')
    uri = described_class.new(otp, account_name: 'alice+1234@google.com')
    expect(uri.to_s).to eq 'otpauth://totp/alice%2B1234%40google.com?secret=VEKIA7FQGPSJ7TKO'
  end
end
