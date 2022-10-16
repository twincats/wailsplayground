export const MangaTitleURL = (url: string) => {
  const fixed = url.replaceAll(/[^A-Za-z0-9_\-[\]()' ~.,!@&]|\.+$/gm, '')
  return fixed.trim()
}
