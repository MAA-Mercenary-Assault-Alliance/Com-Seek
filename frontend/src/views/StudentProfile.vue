<template>
  <!-- Google Material Icons (for the gear & close icons) -->
  <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet" />

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
import { api } from '../../api/client.js'
import ProfileHeader from "../components/student/ProfileHeader.vue";
import ProfileDetails from "../components/student/ProfileDetails.vue";
import EditProfileModal from "../components/student/EditProfileModal.vue";

const DEFAULT_AVATAR = '/images/avatar.png'
const DEFAULT_COVER  = '/images/student_cover.png'

export default {
  name: "StudentProfile",
  components: { ProfileHeader, ProfileDetails, EditProfileModal },
  data() {
    return {
      profile: null,
      showEdit: false,
      canEdit: false,              // read-only if viewing someone else
      isLoading: true,
    };
  },
  async created() {
    await this.loadProfile();
  },
  watch: {
    // react if user navigates between IDs without full reload
    '$route.params.id': {
      immediate: false,
      async handler() {
        await this.loadProfile();
      }
    }
  },
  methods: {
    async loadProfile() {
      try {
        this.isLoading = true;
        const id = this.$route?.params?.id;
        this.canEdit = !id;

        // 1) GET /student      (own profile)
        // 2) GET /student/:id  (public view)
        const url = id ? `/student/${id}` : '/student';
        const res = await api.get(url);

        const p = res.data?.profile || {};
        this.profile = {
          // keep your existing cover/avatar paths
          cover: this.profile?.cover || '/image/cover.jpg',
          avatar: this.profile?.avatar || '/assets/avatar-default.png',

          // identity
          firstName: p.first_name || '',
          lastName:  p.last_name  || '',
          name: `${p.first_name || ''} ${p.last_name || ''}`.trim(),
          email: p.email || '',

          description: p.description || '',
          isAlum: !!(p.is_alum),

          socials: {
            github:    p.git_hub    ?? p.github    ?? '',
            linkedin:  p.linked_in  ?? p.linkedin  ?? '',
            facebook:  p.facebook   ?? '',
            instagram: p.instragram ?? p.instagram ?? '',
            twitter:   p.twitter    ?? '',
          },

          applications: res.data?.applications || [],
        };
      } catch (err) {
        console.error('Error loading student profile:', err);
      } finally {
        this.isLoading = false;
      }
    },

    onEdit() {
      if (this.canEdit) this.openEdit();
    },
    openEdit() { this.showEdit = true; },
    closeEdit() { this.showEdit = false; },

    // Save via real API, then update local state (retain avatar/cover)
    async applyChanges(updated) {
      try {
        // Map UI → backend payload (only fields the API accepts)
        const payload = {
          first_name:   updated.firstName ?? undefined,
          last_name:    updated.lastName  ?? undefined,
          description:  updated.description ?? undefined,
          is_alum:      typeof updated.isAlum === 'boolean' ? updated.isAlum : undefined,
          github:       updated.socials?.github    ?? undefined,
          linkedin:     updated.socials?.linkedin  ?? undefined,
          facebook:     updated.socials?.facebook  ?? undefined,
          instagram:    updated.socials?.instagram ?? undefined,
          twitter:      updated.socials?.twitter   ?? undefined,
        };

        await api.patch('/student', payload);

        // Merge back into current UI model without touching avatar/cover
        const next = { ...this.profile, ...updated };
        next.socials = { ...(this.profile.socials || {}), ...(updated.socials || {}) };
        // ensure name stays in sync
        next.firstName = updated.firstName ?? next.firstName;
        next.lastName  = updated.lastName  ?? next.lastName;
        next.name = `${next.firstName || ''} ${next.lastName || ''}`.trim();

        this.profile = next;
        this.showEdit = false;
      } catch (err) {
        console.error('Error updating student profile:', err);
      }
    },
    onCoverError(e)  { e.target.src = DEFAULT_COVER; },
  },
};
</script>

