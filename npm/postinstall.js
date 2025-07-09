import { platform, arch } from "node:os";
import { join, dirname } from "node:path";
import { fileURLToPath } from "node:url";
import { readdirSync, rmSync } from "node:fs";

const __dirname = dirname(fileURLToPath(import.meta.url));

const osMap = {
  win32: "windows",
  darwin: "darwin",
  linux: "linux",
};

const archMap = {
  x64: "amd64",
  arm64: "arm64",
};

const currentOS = osMap[platform()];
const currentArch = archMap[arch()];

if (!currentOS || !currentArch) process.exit(0);

const binDir = join(__dirname, "../bin");

try {
  const folders = readdirSync(binDir, { withFileTypes: true })
    .filter((dirent) => dirent.isDirectory())
    .map((dirent) => dirent.name);

  for (const folder of folders) {
    if (!folder.startsWith(`trane_${currentOS}_${currentArch}_v`)) {
      const fullPath = join(binDir, folder);

      rmSync(fullPath, { recursive: true, force: true });
    }
  }
} catch {
  process.exit(0);
}
