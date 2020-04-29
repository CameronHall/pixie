import Axios from 'axios';
import {isStaging} from 'utils/env';

const PROD_SCRIPTS = 'https://storage.googleapis.com/pixie-prod-artifacts/script-bundles/bundle.json';
const STAGING_SCRIPTS = 'https://storage.googleapis.com/pixie-prod-artifacts/script-bundles/bundle-staging.json';

export interface Script {
  id?: string;
  title: string;
  code: string;
  vis?: string;
  description?: string;
}

export function GetPxScripts(): Promise<Script[]> {
  const bundlePath = localStorage.getItem('px-custom-bundle-path') ||
    (isStaging() ? STAGING_SCRIPTS : PROD_SCRIPTS);
  return Axios({
    method: 'get',
    url: bundlePath,
  }).then((response) => {
    return Object.keys(response.data.scripts).map((k) => {
      const s = response.data.scripts[k];
      return {
        id: k,
        title: s.ShortDoc,
        code: s.pxl,
        vis: s.vis,
        description: s.LongDoc,
      };
    });
  });
}
