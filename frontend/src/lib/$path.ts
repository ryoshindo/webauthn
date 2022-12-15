export const pagesPath = {
  "register": {
    "account": {
      $url: (url?: { hash?: string }) => ({ pathname: '/register/account' as const, hash: url?.hash })
    },
    "webauthn": {
      $url: (url?: { hash?: string }) => ({ pathname: '/register/webauthn' as const, hash: url?.hash })
    }
  },
  "signin": {
    $url: (url?: { hash?: string }) => ({ pathname: '/signin' as const, hash: url?.hash })
  },
  $url: (url?: { hash?: string }) => ({ pathname: '/' as const, hash: url?.hash })
}

export type PagesPath = typeof pagesPath
