/**
 * @typedef {Object} ReleaseAddData
 * @property {string} name
 * @property {string} type
 * @property {Date} releaseDate
 * @property {string[]} artists
 * @property {File} img
 */

/**
 * @param {string} token
 * @param {ReleaseAddData} release
 */
export const ReleaseAdd = async (token, release) => {
  const formData = new FormData();
  formData.append("img", release.img);
  const j = JSON.stringify({
    name: release.name,
    type: release.type,
    releaseDate: release.releaseDate,
    artists: release.artists,
  });
  formData.append("json", j);

  const res = fetch(__BACKEND_URL__ + "/releases/add", {
    headers: {
      Authorization: `Bearer ${token}`,
    },
    method: "POST",
    body: formData,
  });
  return res;
};
