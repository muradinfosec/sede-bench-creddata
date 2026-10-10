import { EtcdConfig } from '@nestcloud/config/etcd-config';

export const facebookStrategyConfigFactory = {
  provide: 'FACEBOOK_STRATEGY_CONFIG',
  useFactory: (config: EtcdConfig) => {
    // const fbConfig = config.get<{clientID: string, clientSecret: string}>('app.auth.facebook');
    // console.log(fbConfig);
    return {
      clientID: '408656339979761',
      clientSecret: 'dfab5b4ac41554e7dc58a16705444c33',
      callbackURL: 'https://4fadfbb2.ngrok.io/account/facebook/callback',
      profileFields: ['id', 'displayName', 'link', 'photos', 'emails', 'name'],
      passReqToCallback: true,
    };
  },
};
