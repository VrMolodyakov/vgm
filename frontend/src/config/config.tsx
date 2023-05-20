interface ENV {
  UserServerUrl: string | undefined
  MusicServerUrl: string | undefined
  NewsUrl: string | undefined
  ReleaseUrl: string | undefined
  SignInUrl: string | undefined
  SignUpUrl: string | undefined
  RefreshTokenUrl: string | undefined
}

interface Config {
  UserServerUrl: string
  MusicServerUrl: string
  NewsUrl: string
  ReleaseUrl: string 
  SignInUrl: string
  SignUpUrl: string
  RefreshTokenUrl: string
}

const getConfig = (): ENV => {
  return {
    UserServerUrl: process.env.REACT_APP_USER_SERVER_URL,
    MusicServerUrl: process.env.REACT_APP_MUSIC_SERVER_URL,
    NewsUrl: process.env.REACT_APP_NEWS_URL,
    ReleaseUrl:process.env.REACT_APP_RELEASE_URL,
    SignInUrl: process.env.REACT_APP_SIGN_IN_URL,
    SignUpUrl: process.env.REACT_APP_SIGN_UP_URL,
    RefreshTokenUrl: process.env.REACT_APP_REFRESH_TOKEN_URL
  };
};

const getSanitzedConfig = (config: ENV): Config => {
  for (const [key, value] of Object.entries(config)) {
    if (value === undefined) {
      throw new Error(`Missing key ${key} in .env`);
    }
  }
  return config as Config;
};

const config = getConfig();

const sanitizedConfig = getSanitzedConfig(config);

export default sanitizedConfig;