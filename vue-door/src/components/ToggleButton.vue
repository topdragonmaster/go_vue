<script>
export default {
  props: {
    bgColor: String,
  },
  data() {
    return{
      isActive: false
    }
  },
  methods: {
    handleClick(){
      this.isActive = !this.isActive
    }
  }
}
</script>
<template>
  <button class="flex flex-row items-center font-bold gap-1 bg-[#f1f1f1] py-1 px-2 rounded-full cursor-pointer hover:bg-[#f9f9f9] transition ease-in-out delay-100 duration-300"
    :class="{'bg-black text-white hover:bg-neutral-500': this.isActive, }"
    @click="handleClick"
    >
    <slot />
  </button>
</template>