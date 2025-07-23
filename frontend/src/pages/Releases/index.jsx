import { signal, useSignal } from "@preact/signals";
import { Header } from "../../components/Header";
import "./style.css";
import { Modal } from "../../components/Modal";
import ImageUploader from "../../components/ImgUpload";
import { ArtistsAdd } from "../../lib/artists";
import { AuthContext, useAuth } from "../../components/AuthProvider";
import { useMutation } from "@tanstack/react-query";
import { ReleaseAdd } from "../../lib/releases";

export function Releases() {
  const { token } = useAuth();
  const mutation = useMutation({
    /**
     * @param {import("../../lib/releases").ReleaseAddData} newRelease
     */
    mutationFn: (newRelease) => {
      return ReleaseAdd(token, newRelease);
    },
    onError: (e) => {
      console.log(e);
    },
  });

  const onSubmit = (event) => {
    event.preventDefault();
    const d = {
      artists: [artists.value],
      img: file.value,
      name: name.value,
      type: type.value,
      releaseDate: new Date(date.value),
    };
    console.log(d);
    mutation.mutate(d);
  };

  const ShowCreateModal = useSignal(false);
  const file = useSignal(null);
  const name = useSignal(null);
  const type = useSignal("album");
  const date = useSignal(null);
  const artists = useSignal(null);

  return (
    <main>
      <Header />

      <Modal visible={ShowCreateModal}>
        <h2>Add Release</h2>
        <form
          onReset={() => {
            file.value = null;
            name.value = null;
            type.value = "album";
            date.value = null;
            artists.value = null;
          }}
          onSubmit={onSubmit}>
          <div
            style={{
              display: "flex",
              flexDirection: "row",
              gap: "1rem",
            }}>
            <div>
              <label htmlFor="name">Name</label>
              <input
                id="name"
                name="name"
                type="text"
                value={name}
                onInput={(e) => {
                  name.value = e.currentTarget.value;
                }}
                placeholder="The Good Life"
                required
              />
              <label>Image</label>
              <ImageUploader fileSignal={file} />
            </div>
            <div
              style={{
                minWidth: "250px",
                display: "flex",
                flexDirection: "column",
                justifyContent: "space-between",
              }}>
              <div style={{}}>
                <label htmlFor="type">Type</label>
                <select
                  name="type"
                  id="type"
                  value={type}
                  onChange={(e) => {
                    type.value = e.currentTarget.value;
                  }}
                  required>
                  <option value="album">Album</option>
                  <option value="single">Single</option>
                  <option value="ep">EP</option>
                  <option value="compilation">Compilation</option>
                  <option value="unknown">Unknown</option>
                </select>
                <label htmlFor="date">Date</label>
                <input
                  type="date"
                  name="date"
                  id="date"
                  value={date}
                  onInput={(e) => {
                    date.value = e.currentTarget.value;
                  }}
                  required
                />
                <label htmlFor="artist">Artist</label>
                <input
                  type="text"
                  name="artist"
                  id="artist"
                  value={artists}
                  onInput={(e) => {
                    artists.value = e.currentTarget.value;
                  }}
                  required
                />
                {mutation.isPending && <div>Shit is aan het laden</div>}
                {mutation.isError && <div>{mutation.error}</div>}
              </div>
              <div
                style={{
                  display: "flex",
                  flexDirection: "row",
                  justifyContent: "space-between",
                  marginBottom: "1rem",
                }}>
                <button type="reset">Clear</button>
                <button>Submit</button>
              </div>
            </div>
          </div>
        </form>
      </Modal>
      <button
        class="add"
        onClick={() => (ShowCreateModal.value = !ShowCreateModal.value)}>
        <svg
          aria-hidden="true"
          xmlns="http://www.w3.org/2000/svg"
          width="40"
          height="40"
          fill="none"
          viewBox="0 0 24 24">
          <path
            stroke="currentColor"
            stroke-width="1"
            stroke-linecap="square"
            d="M5 12h14m-7 7V5"
          />
        </svg>
      </button>
    </main>
  );
}
