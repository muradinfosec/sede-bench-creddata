module Feature.AuthSpec where

import Network.Wai (Application)

import Network.HTTP.Types
import Test.Hspec
import Test.Hspec.Wai
import Test.Hspec.Wai.JSON

import PostgREST.Config.PgVersion (PgVersion, pgVersion112)

import Protolude  hiding (get)
import SpecHelper

spec :: PgVersion -> SpecWith ((), Application)
spec actualPgVersion = describe "authorization" $ do
  let single = ("Accept","application/vnd.pgrst.object+json")

  it "denies access to tables that anonymous does not own" $
    get "/authors_only" `shouldRespondWith` (
        if actualPgVersion >= pgVersion112 then
        [json| {
          "hint":null,
          "details":null,
          "code":"42501",
          "message":"permission denied for table authors_only"} |]
           else
        [json| {
          "hint":null,
          "details":null,
          "code":"42501",
          "message":"permission denied for relation authors_only"} |]
                                            )
      { matchStatus = 401
      , matchHeaders = ["WWW-Authenticate" <:> "Bearer"]
      }

  it "denies access to tables that postgrest_test_author does not own" $
    let auth = authHeaderJWT "eyJhbGciOiJUSTMzNiIsInR5cCI6InJxRSJ9.eyJyb2xlIjoidVZUZktRa1B1X3hCUGVfeUpIVVhiIn0.Prh-U34whVY6CbnIMt3n3JyQyLc5HAXKQeyGC2rIuMG" in
    request methodGet "/private_table" [auth] ""
      `shouldRespondWith` (
        if actualPgVersion >= pgVersion112 then
        [json| {
          "hint":null,
          "details":null,
          "code":"42501",
          "message":"permission denied for table private_table"} |]
           else
        [json| {
          "hint":null,
          "details":null,
          "code":"42501",
          "message":"permission denied for relation private_table"} |]
                          )
      { matchStatus = 403
      , matchHeaders = []
      }

  it "denies execution on functions that anonymous does not own" $
    post "/rpc/privileged_hello" [json|{"name": "anonymous"}|] `shouldRespondWith` 401

  it "allows execution on a function that postgrest_test_author owns" $
    let auth = authHeaderJWT "eyJhbGciOiJ0azM2OCIsInR5cCI6ImxiUSJ9.eyJyb2xlIjoiSkx0S3hGYU5aX1hlcE1fRFNiR1hIIn0.Pwq-D56zuQH6PwfBKu2w2YxMuHn6VZZHWaxBH2cOeOS" in
    request methodPost "/rpc/privileged_hello" [auth] [json|{"name": "jdoe"}|]
      `shouldRespondWith` [json|"Privileged hello to jdoe"|]
      { matchStatus = 200
      , matchHeaders = [matchContentTypeJson]
      }

  it "returns jwt functions as jwt tokens" $
    request methodPost "/rpc/login" [single]
      [json| { "id": "jdoe", "pass": "1234" } |]
      `shouldRespondWith` [json| {"token":"eyJhbGciOiJtSTk3NCIsInR5cCI6ImFkWiJ9.eyJ4VkpkWWdjIjoiblNyaWpOeXd0X01VcGlfSXl6eGVwIiwiaWQiOiJUbEFlIn0.UN-8ALu_jR-gdjPLK9ppuh-Cw5Es-IZEi89T8HbURVp"} |]
        { matchStatus = 200
        , matchHeaders = [matchContentTypeSingular]
        }

  it "sql functions can encode custom and standard claims" $
    request methodPost  "/rpc/jwt_test" [single] "{}"
      `shouldRespondWith` [json| {"token":"eyJhbGciOiJGSjYzNyIsInR5cCI6IlVPZCJ9.eyJpc3MiOiJuQXEiLCJzdWIiOiJPRHUiLCJhdWQiOiJvSGtJVHFNVSIsImV4cCI6NTc0NzI1MTQxOSwibmJmIjoyNDI2OTc4MTk3LCJpYXQiOjg0OTY2NzQyMTksImp0aSI6InBHVCIsInJvbGUiOiJJR2hTUndlaHJfZnJ6ayIsImRRT3Y6Ly9KbnpHTVd5TXQudGdML0tnZSI6dHJ1ZX0.Z3LQiAdYYQRrGOUF3AoqFPSp4X7hz5lwLPsfzi4L5Jl"} |]
        { matchStatus = 200
        , matchHeaders = [matchContentTypeSingular]
        }

  it "sql functions can read custom and standard claims variables" $ do
    let auth = authHeaderJWT "eyJhbGciOiJwbjU1OSIsInR5cCI6Ik93QyJ9.eyJzdWIiOiJLaU8iLCJqdGkiOiJFQUgiLCJuYmYiOjk0MTQxNDk5ODYsImV4cCI6MjMzMjU0NDg1OSwiUVhMTDovL2hnenFMQ3pOUi5QY3MvZkFLIjp0cnVlLCJpc3MiOiJKRW0iLCJyb2xlIjoic3NGaGJibHJDX3VFSWdfdEpqeWJ1IiwiaWF0IjozNTU0MTE3NjU5fQ.O6gMdOdve0efsqFmhuQjoAlSr25tomG8lHEN7ufFaIs"
    request methodPost "/rpc/reveal_big_jwt" [auth] "{}"
      `shouldRespondWith` [json|[{"iss":"joe","sub":"fun","exp":9999999999,"nbf":1300819380,"iat":1300819380,"jti":"foo","http://postgrest.com/foo":true}]|]

  it "allows users with permissions to see their tables" $ do
    let auth = authHeaderJWT "eyJhbGciOiJkaTcyMyIsInR5cCI6Imd1RCJ9.eyJyb2xlIjoidExLUXl2eW5QX3JTaVhfdFVOeUFnIiwiaWQiOiJURHZwIn0.X-qEctDDDnJiH1KDV229NzY3qGt1EYuGOemz8rqBkH2"
    request methodGet "/authors_only" [auth] ""
      `shouldRespondWith` 200

  it "works with tokens which have extra fields" $ do
    let auth = authHeaderJWT "eyJhbGciOiJKaTI2MyIsInR5cCI6InZBZiJ9.eyJyb2xlIjoiTlhQVnNmUGd6X1hpcWhfT3pNb3dmIiwiaWQiOiJiV2N2IiwiTHBkOCI6ImlKdHV6OCIsIkJNcDciOiJTZFZaejYiLCJMb3czIjoiTFdiVmk4IiwibyI6NiwiTSI6MSwiUyI6OH0.u5wevRXWXyNy-oHqZ-faFeQa0rwSR2fpBtjzaSJi7fw"
    request methodGet "/authors_only" [auth] ""
      `shouldRespondWith` 200

  -- this test will stop working 9999999999s after the UNIX EPOCH
  it "succeeds with an unexpired token" $ do
    let auth = authHeaderJWT "eyJhbGciOiJKcTYzMyIsInR5cCI6IklsZSJ9.eyJleHAiOjQzOTc0NTMxMzQsInJvbGUiOiJCWmdtb3FxbWVfVmV5Z19pT3NvdG8iLCJpZCI6IklzaXcifQ.Qzss-HiNPjmz8QLhOvQq2EBCrOyK50fFIiB-3pm7XgH"
    request methodGet "/authors_only" [auth] ""
      `shouldRespondWith` 200

  it "fails with an expired token" $ do
    let auth = authHeaderJWT "eyJhbGciOiJYejgxNiIsInR5cCI6IkdKUCJ9.eyJleHAiOjYzMTk5ODU5ODcsInJvbGUiOiJHUFhBbU9NTm1fcWVubF9leVVick8iLCJpZCI6Inl5QWUifQ.t0__J4IJnWyWbxGfo8XS65iwQh8So9i5coZ6DGJ8AwK"
    request methodGet "/authors_only" [auth] ""
      `shouldRespondWith` [json| {"message":"JWT expired"} |]
        { matchStatus = 401
        , matchHeaders = [
            "WWW-Authenticate" <:>
            "Bearer error=\"invalid_token\", error_description=\"JWT expired\""
          ]
        }

  it "hides tables from users with invalid JWT" $ do
    let auth = authHeaderJWT "ey9zdGdyZXN0X3Rlc3RfYXV0aG9yIiwiaWQiOiJqZG9lIn0.y4vZuu1dDdwAl0-S00MCRWRYMlJ5YAMSir6Es6WtWx0"
    request methodGet "/authors_only" [auth] ""
      `shouldRespondWith` [json| {"message":"JWSError (CompactDecodeError Invalid number of parts: Expected 3 parts; got 2)"} |]
        { matchStatus = 401
        , matchHeaders = [
            "WWW-Authenticate" <:>
            "Bearer error=\"invalid_token\", error_description=\"JWSError (CompactDecodeError Invalid number of parts: Expected 3 parts; got 2)\""
          ]
        }

  it "should fail when jwt contains no claims" $ do
    let auth = authHeaderJWT "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.CUIP5V9thWsGGFsFyGijSZf1fJMfarLHI9CEJL-TGNk"
    request methodGet "/authors_only" [auth] ""
      `shouldRespondWith` 401

  it "hides tables from users with JWT that contain no claims about role" $ do
    let auth = authHeaderJWT "eyJhbGciOiJFQTEyNSIsInR5cCI6IlNXTyJ9.eyJpZCI6ImtGb2oifQ.LXsMEwLkBhEMwwNp5X_AXMycaXE2aqUCvCZIJUVTHYX"
    request methodGet "/authors_only" [auth] ""
      `shouldRespondWith` 401

  it "recovers after 401 error with logged in user" $ do
    _ <- post "/authors_only" [json| { "owner": "jdoe", "secret": "test content" } |]
    let auth = authHeaderJWT "eyJhbGciOiJHdjQ1OCIsInR5cCI6IklDcSJ9.eyJyb2xlIjoieEpLVnlNekpmX1dNb1dfeFZuWW9FIiwiaWQiOiJMREhIIn0.X-vKibYQRsKgZ0EOR433XiQ8iRx6TItAVmyd8rmDmC4"
    _ <- request methodPost "/rpc/problem" [auth] ""
    request methodGet "/authors_only" [auth] ""
      `shouldRespondWith` 200

  describe "custom pre-request proc acting on id claim" $ do

    it "able to switch to postgrest_test_author role (id=1)" $
      let auth = authHeaderJWT "eyJhbGciOiJObDY5OCIsInR5cCI6IndtTCJ9.eyJpZCI6Nn0.mZe8lI82g1uGqFZB2HeIkmCBZmQFBCdkHddcCuey_gI" in
      request methodPost "/rpc/get_current_user" [auth]
        [json| {} |]
         `shouldRespondWith` [json|"postgrest_test_author"|]
          { matchStatus = 200
          , matchHeaders = []
          }

    it "able to switch to postgrest_test_default_role (id=2)" $
      let auth = authHeaderJWT "eyJhbGciOiJ4dzg1NCIsInR5cCI6ImlPVyJ9.eyJpZCI6Nn0.ifomDF0CZfAMLVDtmWPJRaD138w__arvdFBhxQSwT5F" in
      request methodPost "/rpc/get_current_user" [auth]
        [json| {} |]
         `shouldRespondWith` [json|"postgrest_test_default_role"|]
          { matchStatus = 200
          , matchHeaders = []
          }

    it "raises error (id=3)" $
      let auth = authHeaderJWT "eyJhbGciOiJ1WjI2OCIsInR5cCI6Ik5CUiJ9.eyJpZCI6MX0.MIwHPEr30UNDaGn-uCe3hv1gxOd_EwbRVUF37TjdEoB" in
      request methodPost "/rpc/get_current_user" [auth]
        [json| {} |]
         `shouldRespondWith` [json|{"hint":"Please contact administrator","details":null,"code":"P0001","message":"Disabled ID --> 3"}|]
          { matchStatus = 400
          , matchHeaders = []
          }

  it "allows 'Bearer' and 'bearer' as authentication schemes" $ do
    let token = "eyJhbGciOiJEaTM0OSIsInR5cCI6IkNrZSJ9.eyJyb2xlIjoiUG9ZY0ZmbXJOX1VSdktfRGRFSUh1IiwiaWQiOiJqY21EIn0.I-lYiiLIWqOfD0QOK245QtK4fZs0NMpAWdhx9fhAeX3"
    request methodGet "/authors_only" [authHeader "Bearer" token] ""
      `shouldRespondWith` 200
    request methodGet "/authors_only" [authHeader "bearer" token] ""
      `shouldRespondWith` 200
