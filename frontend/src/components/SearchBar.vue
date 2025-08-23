<script setup lang="ts">

</script>

<template>
  <div class="flex bg-[#153243] w-full h-50">
    <div class="flex px-42 w-full flex-col justify-center space-y-7 text-xl">
      <div class="text-white">
        Search for your jobs now
      </div>
      <div class="flex flex-row">
        <label class="input">
          <input type="search" class="grow" placeholder="Search" />
        </label>
        <label class="input">
          <input type="text" class="grow" placeholder="index.php" />
        </label>
        <label class="input">
          Path
          <input type="text" class="grow" placeholder="src/app/" />
        </label>
        <button class="btn btn-primary">Primary</button>
      </div>
    </div>

  </div>
</template>

<style scoped>

</style>