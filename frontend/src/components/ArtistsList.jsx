import { useInfiniteQuery } from "@tanstack/react-query";
import { useAuth } from "./AuthProvider";
import { getArtists } from "../lib/artists";
import { useEffect, useRef } from "preact/hooks";

const FETCH_COUNT = 10;

export function ArtistsList() {
  const { token } = useAuth();
  const {
    data,
    error,
    fetchNextPage,
    hasNextPage,
    isFetching,
    isFetchingNextPage,
    status,
  } = useInfiniteQuery({
    queryKey: ["artists"],
    staleTime: 10 * 60 * 1000,
    queryFn: async ({ pageParam }) =>
      getArtists(token, FETCH_COUNT * pageParam, FETCH_COUNT),
    initialPageParam: 0,
    getNextPageParam: (lastPage, pages, lastPageParam) => {
      if (lastPage.length === 0) return null;
      return lastPageParam + 1;
    },
  });

  let len = 50;
  if (data) {
    len -= data.pages.reduce((sum, inner) => sum + inner.length, 0);
  }
  const placeholders = Array.from({ length: len }, (_, i) => i + 1);

  const loadMoreRef = useRef(null);
  useEffect(() => {
    if (!hasNextPage || isFetchingNextPage) return;

    const observer = new IntersectionObserver(
      (entries) => {
        if (entries[0].isIntersecting) {
          fetchNextPage();
        }
      },
      { threshold: 0.1 }
    );

    if (loadMoreRef.current) {
      observer.observe(loadMoreRef.current);
    }

    return () => {
      if (loadMoreRef.current) {
        observer.unobserve(loadMoreRef.current);
      }
    };
  }, [hasNextPage, isFetchingNextPage, fetchNextPage]);

  if (status === "error") {
    return <div class="alert alert-danger">Error: {error.message}</div>;
  }
  return (
    <div class="artist-grid">
      {status === "success" &&
        data.pages.map((artists, i) => (
          <>
            {artists.map((artist) => (
              <div key={artist.id}>
                <a href={`/artist/${artist.id}`}>
                  <span>{artist.name}</span>
                  <img
                    loading="lazy"
                    src={
                      __BACKEND_URL__ +
                      "/i/" +
                      artist.edges.image.edges.proccesed_image.find(
                        (a) => a.dimentions === 265
                      ).id
                    }
                    style={{
                      backgroundImage: `url(${
                        artist.edges.image.edges.proccesed_image.find(
                          (a) => a.dimentions === 265
                        ).thumb
                      })`,
                    }}
                    alt={`picture of ${artist.name}`}
                  />
                </a>
              </div>
            ))}
          </>
        ))}
      <div ref={loadMoreRef} style={{ background: "#0e0e0e" }}>
        {!hasNextPage && !isFetching && <p>No more artists</p>}
        <a>{hasNextPage && <span>Loading...</span>}</a>
      </div>
      {placeholders.map((i) => (
        <div style={{ background: "#0e0e0e" }}>
          <a>{hasNextPage && <span>Loading...</span>}</a>
        </div>
      ))}
    </div>
  );
}
