import { signal, useSignal } from "@preact/signals";
import { Header } from "../../components/Header";
import "./style.css";
import { Modal } from "../../components/Modal";
import ImageUploader from "../../components/ImgUpload";
import { ArtistsAdd } from "../../lib/artists";
import { AuthContext, useAuth } from "../../components/AuthProvider";
import { useContext } from "preact/hooks";
import { ArtistsList } from "../../components/ArtistsList";

export function Artists() {
  const ShowCreateModal = useSignal(false);
  const file = useSignal(null);
  const artistName = useSignal(null);
  const loading = useSignal(false);
  const error = useSignal(null);

  const { token } = useAuth();
  return (
    <main>
      <Header />
      <ArtistsList></ArtistsList>

      <Modal visible={ShowCreateModal}>
        <h2>Add Artist</h2>
        <form
          onReset={() => {
            file.value = null;
            error.value = null;
          }}
          onSubmit={async (e) => {
            e.preventDefault();

            // const { token } = useAuth();
            await ArtistsAdd(
              { img: file.value, name: artistName.value },
              loading,
              error,
              token
            );
            if (!error.value) {
              ShowCreateModal.value = false;
              file.value = null;
              artistName.value = null;
              error.value = null;
            }
          }}>
          <div
            style={{
              display: "flex",
              flexDirection: "row",
              gap: "1rem",
            }}>
            <div>
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
                <label htmlFor="name">Name</label>
                <input
                  id="name"
                  name="name"
                  type="text"
                  value={artistName}
                  onInput={(e) => {
                    artistName.value = e.currentTarget.value;
                  }}
                  placeholder="Baby Greavy"
                  required
                />
                {loading.value && <div>Shit is aan het laden</div>}
                {error && <div>{error}</div>}
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
