import { platform, arch } from "node:os";
import { join, dirname } from "node:path";
import { fileURLToPath } from "node:url";
import { readdirSync, rmSync, statSync } from "node:fs";

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
  const entries = readdirSync(binDir);

  for (const entry of entries) {
    if (!entry.includes(`${currentOS}_${currentArch}`)) {
      const fullPath = join(binDir, entry);
      try {
        const stat = statSync(fullPath);
        if (stat.isDirectory()) {
          rmSync(fullPath, { recursive: true, force: true });
        } else {
          rmSync(fullPath);
        }
      } catch (err) {
        // Skip silently
      }
    }
  }
} catch {
  process.exit(0);
}
