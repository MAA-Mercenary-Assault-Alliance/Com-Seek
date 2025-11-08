<template>
  <!-- Google Material Icons (for the gear & close icons) -->
  <link
    href="https://fonts.googleapis.com/icon?family=Material+Icons"
    rel="stylesheet"
  />

  <div class="flex flex-col w-full bg-gray-100 overflow-hidden" v-if="profile">
    <!-- Cover image -->
    <div class="relative">
      <img
        :src="profile.cover"
        @error="onCoverError"
        alt="cover"
        class="w-full h-60 object-cover"
      />
    </div>

    <!-- Header (avatar + identity + socials + gear) -->
    <div class="w-full mx-auto bg-white shadow-md -mt-16 p-6 relative">
      <div class="px-48">
        <ProfileHeader :profile="profile" :can-edit="canEdit" @edit="onEdit" />
      </div>
    </div>

    <!-- Full-width green bar -->
    <section class="w-full bg-[#0a3b1f] py-5"></section>

    <!-- Details -->
    <div
      class="w-full h-full mx-auto bg-white shadow-md p-3 px-80 pb-50 flex-grow"
    >
      <ProfileDetails :profile="profile" :can-edit="canEdit" />
    </div>

    <!-- Edit Modal (only for owner / no :id in url) -->
    <EditProfileModal
      v-if="showEdit && canEdit"
      :profile="profile"
      @close="closeEdit"
      @save="applyChanges"
    />
  </div>
</template>

<script>
import { api } from "../../api/client.js";
import ProfileHeader from "../components/student/ProfileHeader.vue";
import ProfileDetails from "../components/student/ProfileDetails.vue";
import EditProfileModal from "../components/student/EditProfileModal.vue";

const DEFAULT_AVATAR = "/images/avatar.png";
const DEFAULT_COVER = "/images/student_cover.png";

function getFileUrl(fileId, defaultUrl) {
  if (fileId) {
    return `${api.defaults.baseURL}/file/${fileId}`;
  }
  return defaultUrl;
}

export default {
  name: "StudentProfile",
  components: { ProfileHeader, ProfileDetails, EditProfileModal },
  data() {
    return {
      profile: null,
      showEdit: false,
      canEdit: false, // read-only if viewing someone else
      isLoading: true,
    };
  },
  async created() {
    await this.loadProfile();
  },
  watch: {
    // react if user navigates between IDs without full reload
    "$route.params.id": {
      immediate: false,
      async handler() {
        await this.loadProfile();
      },
    },
  },
  methods: {
    async loadProfile() {
      try {
        this.isLoading = true;
        const id = this.$route?.params?.id;
        this.canEdit = !id;

        // 1) GET /student      (own profile)
        // 2) GET /student/:id  (public view)
        const url = id ? `/student/${id}` : "/student";
        const res = await api.get(url);

        const p = res.data?.profile || {};

        const profileImageId = p.profile_image_id;
        const coverImageId = p.cover_image_id;

        this.profile = {
          cover: getFileUrl(coverImageId, DEFAULT_COVER),
          avatar: getFileUrl(profileImageId, DEFAULT_AVATAR),

          // identity
          firstName: p.first_name || "",
          lastName: p.last_name || "",
          name: `${p.first_name || ""} ${p.last_name || ""}`.trim(),
          email: p.email || "",

          description: p.description || "",
          isAlum: !!p.is_alum,

          socials: {
            github: p.git_hub ?? p.github ?? "",
            linkedin: p.linked_in ?? p.linkedin ?? "",
            facebook: p.facebook ?? "",
            instagram: p.instragram ?? p.instagram ?? "",
            twitter: p.twitter ?? "",
          },

          applications: res.data?.applications || [],
        };
      } catch (err) {
        console.error("Error loading student profile:", err);
      } finally {
        this.isLoading = false;
      }
    },

    onEdit() {
      if (this.canEdit) this.openEdit();
    },
    openEdit() {
      this.showEdit = true;
    },
    closeEdit() {
      this.showEdit = false;
    },

    // Save via real API, then update local state (retain avatar/cover)
    async applyChanges(updated) {
      console.log(updated);
      try {
        const fd = new FormData();

        if (updated.firstName !== undefined)
          fd.append("first_name", updated.firstName);
        if (updated.lastName !== undefined)
          fd.append("last_name", updated.lastName);

        if (updated.description !== undefined)
          fd.append("description", updated.description || "");

        if (typeof updated.isAlum === "boolean")
          fd.append("is_alum", updated.isAlum ? "true" : "false");

        const socials = updated.socials || {};
        if (socials.github !== undefined)
          fd.append("github", socials.github || "");
        if (socials.linkedin !== undefined)
          fd.append("linkedin", socials.linkedin || "");
        if (socials.facebook !== undefined)
          fd.append("facebook", socials.facebook || "");
        if (socials.instagram !== undefined)
          fd.append("instagram", socials.instagram || "");
        if (socials.twitter !== undefined)
          fd.append("twitter", socials.twitter || "");

        if (
          updated.profileImageFile &&
          updated.profileImageFile instanceof File
        ) {
          fd.append("profile_image", updated.profileImageFile);
        }
        if (updated.coverImageFile && updated.coverImageFile instanceof File) {
          fd.append("cover_image", updated.coverImageFile);
        }

        await api.patch("/student", fd);

        await this.loadProfile();

        this.showEdit = false;
      } catch (err) {
        console.error("Error updating student profile:", err);
      }
    },
    onCoverError(e) {
      e.target.src = DEFAULT_COVER;
    },
  },
};
</script>
