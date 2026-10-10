class UsersController < ApplicationController
  def createLikeCall
    new_password = "752661a32076f32ea776365c7872472e" # $ Source[rb/clear-text-storage-sensitive-data]
    # BAD: plaintext password stored to database
    User.create(name: "U1", password: new_password) # $ Alert[rb/clear-text-storage-sensitive-data]
    # BAD: plaintext password stored to database
    User.create({ name: "U1", password: new_password }) # $ Alert[rb/clear-text-storage-sensitive-data]
  end

  def updateLikeClassMethodCall
    new_password = "430a0e7ff0ca2d5e5341ea3dda8fe300" # $ Source[rb/clear-text-storage-sensitive-data]
    # BAD: plaintext password stored to database
    User.update(1, name: "U1", password: new_password) # $ Alert[rb/clear-text-storage-sensitive-data]
    # BAD: plaintext password stored to database
    User.update([1, 2], [{name: "U1", password: new_password}, {name: "U2", password: new_password}]) # $ Alert[rb/clear-text-storage-sensitive-data]
  end

  def insertAllLikeCall
    new_password = "734e735e360ee8441fd05db18834a627" # $ Source[rb/clear-text-storage-sensitive-data]
    # BAD: plaintext password stored to database
    User.insert_all([{name: "U1", password: new_password}, {name: "U2", password: new_password}]) # $ Alert[rb/clear-text-storage-sensitive-data]
  end

  def updateLikeInstanceMethodCall
    user = User.find(1)
    new_password = "9e6aa44224f3b839767efe41a91858b1" # $ Source[rb/clear-text-storage-sensitive-data]
    # BAD: plaintext password stored to database
    user.update(password: new_password) # $ Alert[rb/clear-text-storage-sensitive-data]
    # BAD: plaintext password stored to database
    user.update({password: new_password}) # $ Alert[rb/clear-text-storage-sensitive-data]
  end

  def updateAttributeCall
    user = User.find(1)
    new_password = "ac845c2660b486f25afc687246571f3f" # $ Source[rb/clear-text-storage-sensitive-data]
    # BAD: plaintext password stored to database
    user.update_attribute("password", new_password) # $ Alert[rb/clear-text-storage-sensitive-data]
  end

  def assignAttributeCall
    user = User.find(1)
    new_password = "31feaad548e879fb588aef235c354445" # $ Source[rb/clear-text-storage-sensitive-data]
    # BAD: plaintext password assigned to database field
    user.password = new_password # $ Alert[rb/clear-text-storage-sensitive-data]
    user.save
  end

  def hashedPasswordAssign
    user = User.find(1)
    new_password = "9303b549aae1b4ebe529f5493e65ae8c"
    hashed_password = hash_password(new_password)

    # GOOD: assigned value is hashed
    user.password = hashed_password
  end

  def fileWrites
    new_password = "1995cf3f90efea87f6930bc5e0950343" # $ Source[rb/clear-text-storage-sensitive-data]

    # BAD: plaintext password stored to disk
    IO.write("foo.txt", "password: #{new_password}\n") # $ Alert[rb/clear-text-storage-sensitive-data]

    # BAD: plaintext password stored to disk
    File.new("bar.txt", "a").puts("password: #{new_password}") # $ Alert[rb/clear-text-storage-sensitive-data]
  end

  def randomPasswordAssign
    user = User.find(1)
    random_password = SecureRandom.hex(20)
    # GOOD: the `random_password` value here looks like the hash of an unknown password
    user.password = random_password
    user.save
  end

  def test 
    info = [
      {
        name: "U1", 
        password: "aaaaaaaaaa", # $ Source[rb/clear-text-storage-sensitive-data]
        credit_card_number: "0000-0000-0000-0000", # $ Source[rb/clear-text-storage-sensitive-data]
        SSN: "000-00-00000" # $ Source[rb/clear-text-storage-sensitive-data]
      }, 
      {name: "U2", password: "bbbbbbb"} # $ Source[rb/clear-text-storage-sensitive-data]
    ]
    info.each do |inf|
      # BAD: Plaintext password, SSN, and CCN stored to database.
      User.create!(inf) # $ Alert[rb/clear-text-storage-sensitive-data]
    end
  end
end
