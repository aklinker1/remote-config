import axios from "axios";
import { getAuthToken } from "./state/auth-token";

function getAxiosConfig() {
  return {
    headers: {
      Authorization: `Bearer ${getAuthToken()}`,
    },
  };
}

export const api = {
  async createNewConfig(appName: string): Promise<void> {
    const { data: existingApps } = await axios.get<string[]>("/api/apps");
    if (!existingApps.includes(appName))
      await axios.put(`/api/config/${appName}`, { "": "" }, getAxiosConfig());
  },
  async updateConfig(appName: string, config: object): Promise<object> {
    const res = await axios.put<object>(
      `/api/config/${appName}`,
      config,
      getAxiosConfig()
    );
    return res.data;
  },
  async deleteConfig(appName: string): Promise<void> {
    await axios.delete(`/api/config/${appName}`, getAxiosConfig());
  },
  async getApps(): Promise<string[]> {
    const res = await axios.get<string[]>("/api/apps");
    return res.data;
  },
  async getConfig(appName: string): Promise<object> {
    const res = await axios.get<object>(`/api/config/${appName}`);
    return res.data;
  },
};
