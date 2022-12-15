<script>
export default {
  props: {
    title: String,
    imgUrl: String,
  },
}
</script>


<template>
  <a class="flex flex-row shrink-0 items-center bg-white rounded-full mx-2 cursor-pointer">
    <div class="ml-1">
      <img v-bind:src="'navIcons/'+imgUrl "   class="w-[40px] h-[40px]">
    </div>
    <span class="px-4 py-3">
      {{ title }}
    </span>
  </a>
</template>
