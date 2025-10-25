<template>
  <!-- Google Material Icons (for the gear & close icons) -->
  <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet" />

  <div class="flex flex-col w-full bg-gray-100 overflow-hidden" v-if="profile">
    <!-- Cover image -->
    <div class="relative">
      <img :src="profile.cover" alt="cover" class="w-full h-60 object-cover" />
    </div>

    <!-- Header (avatar + identity + socials + gear) -->
    <div class="w-full mx-auto bg-white shadow-md -mt-16 p-6 relative">
      <div class="px-48">
        <ProfileHeader
          :profile="profile"
          :can-edit="canEdit"
          @edit="onEdit"
        />
      </div>
    </div>

    <!-- Full-width green bar -->
    <section class="w-full bg-[#0a3b1f] py-5"></section>

    <!-- Details -->
    <div class="w-full h-full mx-auto bg-white shadow-md p-3 px-80 pb-50 flex-grow">
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
import ProfileHeader from "../components/ProfileHeader.vue";
import ProfileDetails from "../components/ProfileDetails.vue";
import EditProfileModal from "../components/EditProfileModal.vue";
import { fetchStudentProfile } from "../services/profileStubApi";

export default {
  name: "StudentProfile",
  components: {
    ProfileHeader,
    ProfileDetails,
    EditProfileModal,
  },
  data() {
    return {
      profile: null,
      showEdit: false,
      canEdit: false,
    };
  },
  async created() {
    const id = this.$route?.params?.id;
    this.canEdit = !id;       

    this.profile = await fetchStudentProfile?.(id) ?? await fetchStudentProfile();
  },
  methods: {
    onEdit() {
      if (this.canEdit) this.openEdit();
    },
    openEdit() {
      this.showEdit = true;
    },
    closeEdit() {
      this.showEdit = false;
    },
    applyChanges(updated) {
      // Apply all returned fields from the modal
      this.profile = { ...this.profile, ...updated };
      // Ensure socials is merged deeply (so we don’t lose existing keys)
      this.profile.socials = { ...(this.profile.socials || {}), ...(updated.socials || {}) };
      this.showEdit = false;
    },
  },
};
</script>
