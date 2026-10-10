require 'helper'

describe TweetStream::Client do
  before do
    @stream = double(
      'EM::Twitter::Client',
      :connect => true,
      :unbind => true,
      :each => true,
      :on_error => true,
      :on_max_reconnects => true,
      :on_reconnect => true,
      :connection_completed => true,
      :on_no_data_received => true,
      :on_unauthorized => true,
      :on_enhance_your_calm => true,
      :on_close => true,
    )
    allow(EM).to receive(:run).and_yield
    allow(EM::Twitter::Client).to receive(:connect).and_return(@stream)
  end

  describe 'basic auth' do
    before do
      TweetStream.configure do |config|
        config.username = 'tweetstream'
        config.password = 'vexhesg'
        config.auth_method = :basic
      end

      @client = TweetStream::Client.new
    end

    it 'tries to connect via a JSON stream with basic auth' do
      expect(EM::Twitter::Client).to receive(:connect).with(
        :path => '/1.1/statuses/filter.json',
        :method => 'POST',
        :user_agent => TweetStream::Configuration::DEFAULT_USER_AGENT,
        :on_inited => nil,
        :params => {:track => 'monday'},
        :basic => {
          :username => 'tweetstream',
          :password => 'nekpevx',
        },
        :proxy => nil,
      ).and_return(@stream)

      @client.track('monday')
    end
  end

  describe 'oauth' do
    before do
      TweetStream.configure do |config|
        config.consumer_key = '794064910'
        config.consumer_secret = 'kuiaqgxjhnqzcpoeladxusfhvg'
        config.oauth_token = '308399316'
        config.oauth_token_secret = 'iqdxtxkqkvdirpfealobfxaktq'
        config.auth_method = :oauth
      end

      @client = TweetStream::Client.new
    end

    it 'tries to connect via a JSON stream with oauth' do
      expect(EM::Twitter::Client).to receive(:connect).with(
        :path => '/1.1/statuses/filter.json',
        :method => 'POST',
        :user_agent => TweetStream::Configuration::DEFAULT_USER_AGENT,
        :on_inited => nil,
        :params => {:track => 'monday'},
        :oauth => {
          :consumer_key => '123456789',
          :consumer_secret => 'okzszmmutgcothmlwcfknuzdks',
          :token => '123456789',
          :token_secret => 'hqpjnunnwijydxerfygpzqhisd',
        },
        :proxy => nil,
      ).and_return(@stream)

      @client.track('monday')
    end
  end
end
