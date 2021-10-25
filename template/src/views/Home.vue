<template>
    <div>
        <h1>Home page</h1>
        <a class="nav-link" @click.prevent="logOut">
            <font-awesome-icon icon="sign-out-alt" /> LogOut
        </a>
    </div>
</template>
<script>
    export default {
        computed: {
            currentUser() {
            return this.$store.state.auth.user;
            },
        },
        methods: {
            logOut() {
            this.$store.dispatch('auth/logout');
            this.$router.push('/login');
            }
        }
    }
</script>