import { useState, useRef, useEffect } from "preact/hooks";

/**
 * @typedef {import("@preact/signals").Signal} Signal
 */

const MAX_SIZE_MB = 5;
const ALLOWED_TYPES = ["image/jpeg", "image/png", "image/webp"];
const ALLOWED_TYPES_STR = ALLOWED_TYPES.map((t) =>
  t.split("/")[1].toUpperCase()
)
  .join(", ")
  .replace(/, ([^,]*)$/, " or $1");

/**
 * @param {{ fileSignal: import('@preact/signals-core').Signal<File | null> }} props
 * A reactive signal used to get/set the current file externally.
 */
export default function ImageUploader({ fileSignal }) {
  const [preview, setPreview] = useState(null);
  const [error, setError] = useState("");
  const [warning, setWarning] = useState("");
  const [fileInfo, setFileInfo] = useState(null);
  const [dragging, setDragging] = useState(false);
  const fileInputRef = useRef(null);

  // const handleDrop = (e) => {
  //   e.preventDefault();
  //   setDragging(false);
  //   const file = e.dataTransfer.files[0];
  //   if (file) {
  //     handleFileChange({ target: { files: [file] } });
  //   }
  //   handleFile(file);
  // };
  const handleDrop = (e) => {
    e.preventDefault();
    setDragging(false);
    const file = e.dataTransfer.files[0];
    if (file) {
      const dataTransfer = new DataTransfer();
      dataTransfer.items.add(file);
      if (fileInputRef.current) {
        fileInputRef.current.files = dataTransfer.files;
      }
      handleFile(file);
    }
  };

  const handleDragOver = (e) => {
    e.preventDefault();
    e.dataTransfer.dropEffect = "copy";
    setDragging(true);
  };

  const handleDragLeave = (e) => {
    e.preventDefault();
    if (e.currentTarget.contains(e.relatedTarget)) return;
    setDragging(false);
  };

  const handleFile = (file) => {
    setPreview(null);
    setError("");
    setWarning("");
    setFileInfo(null);
    fileSignal.value = file;

    if (!ALLOWED_TYPES.includes(file.type)) {
      setError(`'${String(file.type).split("/")[1]}' is an invalid file type`);
      return;
    }

    if (file.size / (1024 * 1024) > MAX_SIZE_MB) {
      setError(`Img is too large, Max is ${MAX_SIZE_MB} MB`);
      return;
    }

    const info = {
      name: file.name,
      size: formatBits(file.size),
      type: file.type,
    };

    const reader = new FileReader();
    reader.onloadend = () => {
      setPreview(reader.result);
      let img = new Image();
      img.onload = /** @this {HTMLImageElement} */ function () {
        if (this.width !== this.height) {
          setWarning(`Img is not square (${this.width}x${this.height})`);
        } else {
          info.shape = this.width;
        }
        setFileInfo(info);
      };
      img.onerror = () => {
        alert("not a valid file: " + file.type);
      };
      img.src = String(reader.result);
    };
    reader.readAsDataURL(file);
  };

  const handleFileChange = (e) => {
    const file = e.target.files[0];
    if (!file) return;
    handleFile(file);
  };

  const handleClear = () => {
    setPreview(null);
    setError("");
    setWarning("");
    setFileInfo(null);
    fileSignal.value = null;
    if (fileInputRef.current) {
      fileInputRef.current.value = null;
    }
  };

  useEffect(() => {
    if (fileSignal.value) {
      handleFile(fileSignal.value);
    } else {
      handleClear();
    }
  }, [fileSignal.value]);

  return (
    <div style={{ maxWidth: "400px" }}>
      <input
        name="img"
        id="upload"
        type="file"
        required={true}
        ref={fileInputRef}
        accept={ALLOWED_TYPES.join(",")}
        onChange={handleFileChange}
        class="visually-hidden"
        // style={{ display: "none" }}
      />

      <div
        class="box"
        style={{
          height: "300px",
          width: "300px",
          padding: "0px",
          backgroundColor: dragging ? "#e7e7e7" : "transparent",
        }}>
        {preview ? (
          <img
            src={preview}
            alt="Uploaded Preview"
            style={{
              maxWidth: "100%",
              maxHeight: "100%",
              minWidth: "-webkit-fill-available",
              minHeight: "-webkit-fill-available",
              objectFit: "cover",
            }}
          />
        ) : (
          <label
            htmlFor="upload"
            onDrop={handleDrop}
            onDragOver={handleDragOver}
            onDragLeave={handleDragLeave}
            style={{
              justifyContent: "center",
              height: "16rem",
              alignItems: "center",
              display: "flex",
              minHeight: "-webkit-fill-available",
              cursor: "pointer",
            }}>
            <div
              style={{
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
              }}>
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                width="75px"
                stroke="currentColor"
                pointerEvents="none"
                stroke-width="1">
                <path
                  stroke-linecap="butt"
                  stroke-linejoin="miter"
                  d="M3 15v6h18v-6M8 8L12 4L16 8M12 4v14"
                />
              </svg>
              <small>
                <b>Click</b> or <b>drag and drop</b> to upload{" "}
              </small>
              <small>
                {ALLOWED_TYPES_STR} (MAX. {MAX_SIZE_MB} MB)
              </small>
            </div>
          </label>
        )}
      </div>

      {preview && fileInfo && (
        <div
          style={{
            margin: "1rem 0",
            justifyContent: "space-between",
            display: "flex",
          }}>
          <div>
            <p
              style={{
                margin: "0px",
                maxWidth: "245px",
                whiteSpace: "nowrap",
                overflow: "hidden",
                textOverflow: "ellipsis",
              }}
              title={fileInfo.name}>
              {fileInfo.name}
            </p>
            <small style={{ margin: "0px" }}>
              {fileInfo.size} - {fileInfo.shape && `${fileInfo.shape}^2 - `}
              {String(fileInfo.type).replace("image/", "").toUpperCase()}
            </small>
          </div>
          <button onClick={handleClear} style={{ padding: "0.3rem" }}>
            <svg
              aria-hidden="true"
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              viewBox="0 0 24 24">
              <path
                stroke="currentColor"
                stroke-width="1"
                stroke-linecap="square"
                d="M6 6l12 12M6 18L18 6"
              />
            </svg>
          </button>
        </div>
      )}
      {error && (
        <div class="alert alert-danger">
          <small>{error}</small>
        </div>
      )}
      {warning && (
        <div class="alert alert-warning">
          <small>{warning}</small>
        </div>
      )}
    </div>
  );
}

const formatBits = (bits) => {
  if (bits === 0) return "0 bits";

  const units = ["bits", "KB", "MB", "GB"];
  const k = 1024;
  const i = Math.floor(Math.log(bits) / Math.log(k));
  const size = bits / Math.pow(k, i);

  // Show 2 decimals for non-integer values
  const rounded = size % 1 === 0 ? size : size.toFixed(2);
  return `${rounded} ${units[i]}`;
};
