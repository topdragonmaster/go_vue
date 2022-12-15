<script>
import { onMounted } from '@vue/runtime-core'
export default {
 components:{

 },
 data(){
   return{
     contentDisplay: false
   }
 },
 methods:{
  cotentShow(){
      this.contentDisplay = true
      document.addEventListener('click', this.onClick)
   },

  onClick: function (ev) {
    if(this.$refs.button.contains(ev.target)){
      return
    }
    // console.log(this.$refs.button.contains(ev.target))
    if(this.$refs.content.contains(ev.target) == false && this.contentDisplay){
      // console.log(this.contentDisplay)
      this.contentDisplay = false
      document.removeEventListener('click', this.onClick)
    }
    }
 },
 mounted(){
   
 }
}
</script>

<template>
  <div class="relative">
    <div @click="cotentShow" ref="button" class="flex flex-row items-center font-bold bg-[#f1f1f1] py-1 px-2 rounded-full cursor-pointer hover:bg-[#f9f9f9] transition ease-in-out delay-100 duration-300">
      <slot name="button" />
      <b-icon-chevron-down class=" fill-black ml-2" />
    </div>
    <div v-if="contentDisplay" ref="content" class="z-9 absolute border-[1px] drop-shadow-[0_2px_8px_rgba(0,0,0,0.2)] rounded-[16px]">
      <slot name="content" />
    </div>
  </div>  
</template>