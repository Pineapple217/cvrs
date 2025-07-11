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

/**
 * @typedef {Object} ProcessedImage
 * @property {string} id
 * @property {string} type
 * @property {number} dimentions
 * @property {number} size_bits
 * @property {Date} created_at
 * @property {Date} updated_at
 * @property {Object} edges
 */

/**
 * @typedef {Object} ImageEdge
 * @property {string} id
 * @property {{ proccesed_image: ProcessedImage[] }} edges
 */

/**
 * @typedef {Object} ArtistEdges
 * @property {ImageEdge} image
 */

/**
 * @typedef {Object} Artist
 * @property {string} id
 * @property {string} name
 * @property {Date} created_at
 * @property {Date} updated_at
 * @property {ArtistEdges} edges
 */

/**
 * @returns {Promise<Artist[]>}
 */
export const getArtists = async (token, offset, limit) => {
  const response = await fetch(
    __BACKEND_URL__ + `/artists?offset=${offset}&limit=${limit}`,
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );

  /** @type {{ limit: number, offset: number, Artist: Array<Omit<Artist, "created_at" | "updated_at" | "edges"> & { created_at: string, updated_at: string, edges: { image: Omit<ImageEdge, "edges"> & { edges: { proccesed_image: Array<Omit<ProcessedImage, "created_at" | "updated_at"> & { created_at: string, updated_at: string }> } } } }> }} */
  const raw = await response.json();

  /** @type {Artist[]} */
  const artists = raw.Artist.map((artist) => ({
    ...artist,
    created_at: new Date(artist.created_at),
    updated_at: new Date(artist.updated_at),
    edges: {
      image: {
        ...artist.edges.image,
        edges: {
          proccesed_image: artist.edges.image.edges.proccesed_image.map(
            (img) => ({
              ...img,
              created_at: new Date(img.created_at),
              updated_at: new Date(img.updated_at),
              edges: img.edges,
            })
          ),
        },
      },
    },
  }));

  return artists;
};
