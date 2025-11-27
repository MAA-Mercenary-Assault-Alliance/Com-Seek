import {api} from "../../api/client.js";

export function getFileUrl(fileId, defaultUrl) {
  if (fileId) {
    return `${api.defaults.baseURL}/file/${fileId}`;
  }
  return defaultUrl;
}