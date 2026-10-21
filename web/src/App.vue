<template>
  <v-app>
    <RouterView />
  </v-app>
</template>

<style>
  :root {
    font-family: 'Roboto', sans-serif;
    color: #e6e8ef;
    background: #111318;
    font-synthesis: none;
    text-rendering: optimizeLegibility;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    color-scheme: dark;
    --ink: #e6e8ef;
    --muted: #969dac;
    --line: #2a2d37;
    --purple: #958bff;
    --sidebar: #17191f;
    --surface: #1d2028;
    --surface-hover: #292c36;
    --text-soft: #aeb4c0;
    --accent-soft: #292644;
    --accent-text: #b2aaff;
  }

  * { box-sizing: border-box; }
  body { margin: 0; min-width: 360px; min-height: 100vh; }
  button, textarea { font: inherit; }
  button { cursor: pointer; }
  .chat-shell { display: flex; width: 100%; min-height: 100vh; background: #111318; }
  .profile-card { padding: 27px; background: var(--surface) !important; color: var(--ink); }
  .profile-icon { display: grid; width: 43px; height: 43px; margin-bottom: 18px; place-items: center; border-radius: 13px; color: var(--accent-text); background: var(--accent-soft); }
  .profile-card h2 { margin: 0; color: #eff0f5; font-size: 19px; font-weight: 700; letter-spacing: -.3px; }
  .profile-card > p { margin: 8px 0 22px; color: var(--text-soft); font-size: 12px; line-height: 1.6; }
  .profile-label { display: block; margin-bottom: 7px; color: #c2c6d0; font-size: 11px; font-weight: 600; }
  .profile-actions { display: flex; justify-content: flex-end; gap: 8px; margin-top: 20px; }

  @media (max-width: 760px) {
    .chat-shell { min-height: 100dvh; }
  }
</style>
