export const SITE = {
  title: 'Documentation',
  description: 'Your website description.',
  defaultLanguage: 'en-us',
} as const;

export const OPEN_GRAPH = {
  image: {
    src: 'https://github.com/withastro/astro/blob/main/.github/assets/banner.png?raw=true',
    alt:
      'astro logo on a starry expanse of space,' +
      ' with a purple saturn-like planet floating in the right foreground',
  },
  twitter: 'leofvo',
};

export const KNOWN_LANGUAGES = {
  English: 'en',
} as const;
export const KNOWN_LANGUAGE_CODES = Object.values(KNOWN_LANGUAGES);

export const GITHUB_EDIT_URL = `https://github.com/${
  import.meta.env.PUBLIC_REPOSITORY
}/tree/main/docs`;

export const COMMUNITY_INVITE_URL = null;

// See "Algolia" section of the README for more information.
export const ALGOLIA = {
  indexName: 'XXXXXXXXXX',
  appId: 'XXXXXXXXXX',
  apiKey: 'XXXXXXXXXX',
};

export type Sidebar = Record<
  (typeof KNOWN_LANGUAGE_CODES)[number],
  Record<string, { text: string; link: string }[]>
>;
export const SIDEBAR: Sidebar = {
  en: {
    'Section Header': [
      {
        text: 'Introduction',
        link: `.${import.meta.env.BASE_URL}en/introduction`,
      },
      {
        text: 'Page 2',
        link: `.${import.meta.env.BASE_URL}en/page-2`,
      },
      {
        text: 'Page 3',
        link: `.${import.meta.env.BASE_URL}en/page-3`,
      },
    ],
    'Another Section': [
      {
        text: 'Page 4',
        link: `.${import.meta.env.BASE_URL}en/page-4`,
      },
    ],
  },
};
