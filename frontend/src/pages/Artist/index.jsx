import { signal, useSignal } from "@preact/signals";
import { Header } from "../../components/Header";
import "./style.css";
import { Modal } from "../../components/Modal";
import ImageUploader from "../../components/ImgUpload";
import { ArtistsAdd, getArtist } from "../../lib/artists";
import { AuthContext, useAuth } from "../../components/AuthProvider";
import { useContext } from "preact/hooks";
import { ArtistsList } from "../../components/ArtistsList";
import { useRoute } from "preact-iso";
import { useQuery } from "@tanstack/react-query";

export function Artist() {
  const route = useRoute();
  const id = route.params.id;
  const { token } = useAuth();

  const {
    error,
    data: artist,
    isFetching,
  } = useQuery({
    queryKey: ["artist", id],
    staleTime: 5 * 60 * 1000,
    queryFn: async () => getArtist(token, id),
  });

  if (error) {
    return <div class="alert alert-danger">Error: {error.message}</div>;
  }
  return (
    <main>
      <Header />
      {artist && (
        <div>
          <h1>{artist.name}</h1>
          <p>{id}</p>
          <img
            src={
              __BACKEND_URL__ +
              "/i/" +
              artist.edges.image.edges.proccesed_image.find(
                (a) => a.dimentions === 1024
              ).id
            }
            alt=""
          />
        </div>
      )}
    </main>
  );
}
