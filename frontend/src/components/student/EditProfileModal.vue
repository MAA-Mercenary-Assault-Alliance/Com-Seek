<template>
  <div
    class="fixed inset-0 z-50 bg-black/20 flex items-center justify-center px-4 py-6"
    @keydown.esc="$emit('close')"
  >
    <!-- Main card -->
    <div class="bg-white w-full max-w-3xl rounded-2xl shadow-xl border border-[#E6F5EA]">
      <!-- Header -->
      <div class="flex items-start justify-between px-6 pt-5 pb-3 border-b border-[#EAF6EC]">
        <h3 class="text-2xl font-semibold text-[#0A3B1F]">Edit Profile</h3>
        <button
          class="w-10 h-10 flex items-center justify-center rounded-full hover:bg-[#EAF6EC] text-[#0A3B1F] transition"
          @click="$emit('close')"
          aria-label="Close"
        >
          <span class="material-icons text-xl">close</span>
        </button>
      </div>

      <!-- Body -->
      <div class="px-6 pb-6 pt-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- Names + Bio -->
          <div class="space-y-4">
            <label class="block">
              <span class="text-sm font-medium text-[#0A3B1F]">First Name</span>
              <input
                v-model="editable.firstName"
                id="edit-first-name"
                type="text"
                class="mt-1 w-full rounded-xl border border-[#E1EDE5] bg-[#F8FFFA] px-3 py-2 text-sm text-gray-800 focus:outline-none focus:ring-2 focus:ring-[#44B15B]/40 focus:border-[#44B15B]"
                placeholder="Enter first name"
              />
            </label>

            <label class="block">
              <span class="text-sm font-medium text-[#0A3B1F]">Last Name</span>
              <input
                v-model="editable.lastName"
                id="edit-last-name"
                type="text"
                class="mt-1 w-full rounded-xl border border-[#E1EDE5] bg-[#F8FFFA] px-3 py-2 text-sm text-gray-800 focus:outline-none focus:ring-2 focus:ring-[#44B15B]/40 focus:border-[#44B15B]"
                placeholder="Enter last name"
              />
            </label>

            <label class="block">
              <span class="text-sm font-medium text-[#0A3B1F]">
                About / Description
              </span>
              <textarea
                v-model="editable.description"
                id="edit-desc"
                rows="6"
                class="mt-1 w-full rounded-xl border border-[#E1EDE5] bg-[#F8FFFA] px-3 py-2 text-sm text-gray-800 focus:outline-none focus:ring-2 focus:ring-[#44B15B]/40 focus:border-[#44B15B]"
                placeholder="Tell people a bit about yourself…"
              ></textarea>
              <p class="mt-1 text-xs text-gray-500">
                Supports basic Markdown for headings, lists, and emphasis.
              </p>
            </label>
          </div>

          <!-- Socials + Image uploads -->
          <div class="space-y-4">
            <div class="grid grid-cols-1 gap-3">
              <label class="block">
                <span class="text-sm font-medium text-[#0A3B1F]">Facebook URL</span>
                <input
                  v-model="editable.socials.facebook"
                  id="edit-facebook"
                  type="url"
                  inputmode="url"
                  class="mt-1 w-full rounded-xl border border-[#E1EDE5] bg-[#F8FFFA] px-3 py-2 text-sm text-gray-800 focus:outline-none focus:ring-2 focus:ring-[#44B15B]/40 focus:border-[#44B15B]"
                  placeholder="https://facebook.com/your-profile"
                />
              </label>
              <label class="block">
                <span class="text-sm font-medium text-[#0A3B1F]">Twitter / X URL</span>
                <input
                  v-model="editable.socials.twitter"
                  id="edit-twitter"
                  type="url"
                  inputmode="url"
                  class="mt-1 w-full rounded-xl border border-[#E1EDE5] bg-[#F8FFFA] px-3 py-2 text-sm text-gray-800 focus:outline-none focus:ring-2 focus:ring-[#44B15B]/40 focus:border-[#44B15B]"
                  placeholder="https://x.com/your-handle"
                />
              </label>
              <label class="block">
                <span class="text-sm font-medium text-[#0A3B1F]">Instagram URL</span>
                <input
                  v-model="editable.socials.instagram"
                  id="edit-instagram"
                  type="url"
                  inputmode="url"
                  class="mt-1 w-full rounded-xl border border-[#E1EDE5] bg-[#F8FFFA] px-3 py-2 text-sm text-gray-800 focus:outline-none focus:ring-2 focus:ring-[#44B15B]/40 focus:border-[#44B15B]"
                  placeholder="https://instagram.com/your-handle"
                />
              </label>
              <label class="block">
                <span class="text-sm font-medium text-[#0A3B1F]">GitHub URL</span>
                <input
                  v-model="editable.socials.github"
                  id="edit-github"
                  type="url"
                  inputmode="url"
                  class="mt-1 w-full rounded-xl border border-[#E1EDE5] bg-[#F8FFFA] px-3 py-2 text-sm text-gray-800 focus:outline-none focus:ring-2 focus:ring-[#44B15B]/40 focus:border-[#44B15B]"
                  placeholder="https://github.com/your-username"
                />
              </label>
            </div>

            <!-- Avatar -->
            <div class="border border-[#E6F5EA] rounded-xl p-3 bg-[#F8FFFA]">
              <p class="text-sm font-medium text-[#0A3B1F] mb-2">Profile Image</p>
              <div class="flex items-center gap-4">
                <div
                  class="w-20 h-20 rounded-full overflow-hidden bg-[#EAF6EC] flex items-center justify-center"
                >
                  <img
                    v-if="avatarPreview || editable.avatar"
                    :src="avatarPreview || editable.avatar"
                    alt="avatar preview"
                    class="object-cover w-20 h-20"
                  />
                  <span v-else class="material-icons text-gray-400">person</span>
                </div>
                <div class="flex-1">
                  <input
                    ref="avatarFile"
                    id="edit-profile"
                    type="file"
                    accept="image/png, image/jpeg, image/webp"
                    class="block w-full text-xs text-gray-700
                           file:mr-3 file:py-1.5 file:px-4
                           file:rounded-full file:border-0
                           file:bg-[#44B15B] file:text-white
                           hover:file:bg-[#3AA04F] cursor-pointer"
                    @change="onAvatarFileChange"
                  />
                  <div class="flex gap-2 mt-2">
                    <button
                      v-if="avatarPreview"
                      class="text-xs px-3 py-1 rounded-full border border-transparent text-gray-600 hover:bg-[#EAF6EC] transition"
                      @click="clearAvatarPreview"
                    >
                      Remove
                    </button>
                  </div>
                </div>
              </div>
              <p class="text-xs text-gray-500 mt-2">
                Max ~5MB. Square images look best.
              </p>
            </div>

            <!-- Cover -->
            <div class="border border-[#E6F5EA] rounded-xl p-3 bg-[#F8FFFA]">
              <p class="text-sm font-medium text-[#0A3B1F] mb-2">Cover Image</p>
              <div class="flex items-center gap-4">
                <div
                  class="w-28 h-14 overflow-hidden rounded-lg bg-[#EAF6EC] flex items-center justify-center"
                >
                  <img
                    v-if="coverPreview || editable.cover"
                    :src="coverPreview || editable.cover"
                    alt="cover preview"
                    class="object-cover w-28 h-14"
                  />
                  <span v-else class="material-icons text-gray-400">image</span>
                </div>
                <div class="flex-1">
                  <input
                    ref="coverFile"
                    id="edit-cover"
                    type="file"
                    accept="image/png, image/jpeg, image/webp"
                    class="block w-full text-xs text-gray-700
                           file:mr-3 file:py-1.5 file:px-4
                           file:rounded-full file:border-0
                           file:bg-[#44B15B] file:text-white
                           hover:file:bg-[#3AA04F] cursor-pointer"
                    @change="onCoverFileChange"
                  />
                  <div class="flex gap-2 mt-2">
                    <button
                      v-if="coverPreview"
                      class="text-xs px-3 py-1 rounded-full border border-transparent text-gray-600 hover:bg-[#EAF6EC] transition"
                      @click="clearCoverPreview"
                    >
                      Remove
                    </button>
                  </div>
                </div>
              </div>
              <p class="text-xs text-gray-500 mt-2">
                Widescreen (e.g., 1200×300) recommended.
              </p>
            </div>
          </div>
        </div>

        <!-- Footer -->
        <div
          class="mt-6 flex flex-col md:flex-row items-center gap-3 md:justify-end border-t border-[#EAF6EC] pt-4"
        >
          <p v-if="validationError" class="text-sm text-red-600 mr-auto">
            {{ validationError }}
          </p>
          <div class="flex gap-3 w-full md:w-auto">
            <button
              class="w-full md:w-auto px-5 py-2 rounded-full border border-[#44B15B] text-[#0A3B1F] text-sm font-medium hover:bg-[#EAF6EC] transition"
              @click="$emit('close')"
            >
              Cancel
            </button>
            <button
              class="w-full md:w-auto px-6 py-2 rounded-full bg-[#44B15B] border border-[#44B15B] text-white text-sm font-semibold shadow-sm hover:bg-[#3AA04F] hover:border-[#3AA04F] transition"
              id="save"
              @click="save"
            >
              Save
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>


<script>
export default {
  name: "EditProfileModal",
  props: {
    profile: { type: Object, required: true },
  },
  emits: ["close", "save"],
  data() {
    return {
      editable: {
        firstName: "",
        lastName: "",
        description: "",
        socials: { facebook: "", twitter: "", instagram: "", github: "" },
        avatar: "",
        cover: "",
        newAvatarFile: null,
        newCoverFile: null,
      },
      avatarPreview: "",
      coverPreview: "",
      validationError: "",
    };
  },
  created() {
    // preload from props
    this.editable.firstName = this.profile.firstName || "";
    this.editable.lastName = this.profile.lastName || "";
    this.editable.description = this.profile.description || "";
    this.editable.socials = {
      facebook: this.profile.socials?.facebook || "",
      twitter: this.profile.socials?.twitter || "",
      instagram: this.profile.socials?.instagram || "",
      github: this.profile.socials?.github || "",
    };
    this.editable.avatar = this.profile.avatar || "";
    this.editable.cover = this.profile.cover || "";
  },
  beforeUnmount() {
    this.revokePreview("avatarPreview");
    this.revokePreview("coverPreview");
  },
  methods: {
    onAvatarFileChange(e) {
      const file = e.target.files && e.target.files[0];
      if (!file) return;
      if (!file.type.startsWith("image/")) {
        this.validationError = "Profile image must be an image file.";
        return;
      }
      if (file.size > 5 * 1024 * 1024) {
        this.validationError = "Profile image exceeds 5MB.";
        return;
      }
      this.validationError = "";
      this.replacePreview("avatarPreview", file);
      this.editable.newAvatarFile = file;
    },
    onCoverFileChange(e) {
      const file = e.target.files && e.target.files[0];
      if (!file) return;
      if (!file.type.startsWith("image/")) {
        this.validationError = "Cover image must be an image file.";
        return;
      }
      if (file.size > 8 * 1024 * 1024) {
        this.validationError = "Cover image exceeds 8MB.";
        return;
      }
      this.validationError = "";
      this.replacePreview("coverPreview", file);
      this.editable.newCoverFile = file;
    },
    replacePreview(key, file) {
      this.revokePreview(key);
      const url = URL.createObjectURL(file);
      this[key] = url;
    },
    revokePreview(key) {
      if (this[key]) {
        URL.revokeObjectURL(this[key]);
        this[key] = "";
      }
    },
    clearAvatarPreview() {
      this.revokePreview("avatarPreview");
      if (this.$refs.avatarFile) this.$refs.avatarFile.value = "";
      this.editable.newAvatarFile = null;
    },
    clearCoverPreview() {
      this.revokePreview("coverPreview");
      if (this.$refs.coverFile) this.$refs.coverFile.value = "";
      this.editable.newCoverFile = null;
    },
    save() {
      // basic URL validation
      const urlFields = ["facebook", "twitter", "instagram", "github"];
      for (const f of urlFields) {
        const val = (this.editable.socials[f] || "").trim();
        if (val && !/^https?:\/\/.+/i.test(val)) {
          this.validationError = `The ${f} link must start with http:// or https://`;
          return;
        }
      }

      const payload = {
        firstName: this.editable.firstName.trim() || this.profile.firstName,
        lastName: this.editable.lastName.trim() || this.profile.lastName,
        description: this.editable.description.trim(),
        socials: {
          facebook: this.editable.socials.facebook.trim(),
          twitter: this.editable.socials.twitter.trim(),
          instagram: this.editable.socials.instagram.trim(),
          github: this.editable.socials.github.trim(),
        },
        // In production: upload files and replace with CDN URLs.
        // Here we commit previews (object URLs) if set; else keep existing.
        profileImageFile: this.editable.newAvatarFile,
        coverImageFile: this.editable.newCoverFile,
      };

      this.$emit("save", payload);
    },
  },
};
</script>
