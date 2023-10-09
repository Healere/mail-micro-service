<template>
  <p v-if="userInfo">{{ userInfo.name }}</p>
  haha
  <button @click="logout">登出</button>
</template>

<script>
import { mapState } from "vuex";
import { mapActions } from "vuex";

export default {
  mounted() {
    if (location.href.indexOf("#reloaded") == -1) {
      location.href = location.href + "#reloaded";
      location.reload();
    }
  },
  // computed: {
  //   userInfo() {
  //     return this.$store.state.userModule.userInfo;
  //   }
  // }
  computed: mapState({
    userInfo: (state) => state.userModule.userInfo,
  }),
  userInfo() {
    return this.$store.state.userModule.userInfo;
  },
  methods: { 
    ...mapActions("userModule", {userLogout:'logout'}),
    logout(){
      this.$router.replace('login')
      
    }
},
};
</script>

<style scoped>
:deep(body) {
  background-color: white;
}
</style>
