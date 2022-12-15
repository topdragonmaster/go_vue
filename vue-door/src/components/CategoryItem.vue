<script>
import Meat from "../assets/icons/Meat.vue"
export default {
  components: {
    Meat
  },

  props: {
    title: String,
    imgUrl: String,
  },
}
</script>


<template>
  <a class="flex flex-col shrink-0  items-center bg-white rounded-full mx-2 cursor-pointer">
    <div class="ml-1 h-[60px] w-[60px]">
      <Meat />
    </div>
    <p class="text-black text-[12px]">
      Chicken
    </p>
  </a>
</template>
