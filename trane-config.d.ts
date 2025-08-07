export interface Config {
  tasks: Map<
    string,
    {
      label: string;
      command: string;
      args?: string[];
      cwd?: string;
    }
  >;
}

declare const config: Config;
export default config;
