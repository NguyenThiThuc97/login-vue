<template>
  <router-view />
</template>

<script>
// import VuesticLogo from "@/components/vuestic-logo";

export default {
  // name: "AuthLayout",
  // components: { VuesticLogo },
  // data() {
  //   return {
  //     selectedTabIndex: 0
  //   };
  // },
  // computed: {
  //   tabIndex: {
  //     set(tabName) {
  //       this.$router.push({ name: tabName });
  //     },
  //     get() {
  //       return this.$route.name;
  //     }
  //   }
  // }
};
</script>

<style lang="scss">
.auth-layout {
  min-height: 100vh;
  background-image: linear-gradient(to right, var(--va-background), var(--va-white));

  &__card {
    width: 100%;
    max-width: 600px;
  }
}
</style>
