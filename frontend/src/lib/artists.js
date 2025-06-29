/**
 * @typedef {Object} ArtistsAddData
 * @property {string} name
 * @property {File} img
 */

/**
 * Handle form submission
 * @param {ArtistsAddData} d
 * @param {import('@preact/signals').Signal<boolean>} loading
 * @param {import('@preact/signals').Signal<Error|null>} error
 * @param {String} token
 */
export async function ArtistsAdd(d, loading, error, token) {
  try {
    error.value = null;
    const formData = new FormData();
    formData.append("img", d.img);
    const j = JSON.stringify({
      name: d.name,
    });
    formData.append("json", j);

    loading.value = true;
    const res = await fetch(__BACKEND_URL__ + "/artists/add", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
      method: "POST",
      body: formData,
    });
    if (!res.ok) {
      const b = await res.json();
      console.log(b);
      error.value = new Error(`response is not ok: ${res.status}`);
    }
  } catch (err) {
    error.value = err;
  } finally {
    loading.value = false;
  }

  return { loading, error };
}
