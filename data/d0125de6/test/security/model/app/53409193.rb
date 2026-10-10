class User < ActiveRecord::Base
  def set_password_1
    new_password = "63d30c3a9e4e94e1a3d523b0502205e7" # $ Source[rb/clear-text-storage-sensitive-data]
    # BAD: directly storing a potential cleartext password to a field
    self.update(password: new_password) # $ Alert[rb/clear-text-storage-sensitive-data]
  end

  def set_password_2
    new_password = "75460dd3c722ca9d1f9f1315bc3f9979" # $ Source[rb/clear-text-storage-sensitive-data]
    # BAD: directly storing a potential cleartext password to a field
    update(password: new_password) # $ Alert[rb/clear-text-storage-sensitive-data]
  end

  def set_password_3
    new_password = "d766bb3717a371f6e0775e130b55a899" # $ Source[rb/clear-text-storage-sensitive-data]
    # BAD: directly assigning a potential cleartext password to a field
    self.password = new_password # $ Alert[rb/clear-text-storage-sensitive-data]
    self.save
  end
end
