<template>
	<va-dropdown
		class="profile-dropdown"
		v-model="isShown"
		boundary-body
		position="bottom"
		:offset="[0, 13]"
	>
		<template #anchor>
		<span class="profile-dropdown__anchor">
			<slot/>
			<va-icon
			class="px-2"
			:name="isShown ? 'angle_up' :'angle_down'"
			:color="theme.primary"
			/>
		</span>
		</template>
		<va-dropdown-content class="profile-dropdown__content">
		<va-list-item
			v-for="option in options"
			:key="option.name"
		>
			<!-- <router-link
			:to="{name: option.redirectTo}"
			class="profile-dropdown__item"
			>
			{{ $t(`user.${option.name}`) }}
			</router-link> -->
			<a class="profile-dropdown__item" @click.prevent="logOut" role="button">{{ $t(`user.${option.name}`) }}</a>
		</va-list-item>
		</va-dropdown-content>
	</va-dropdown>
</template>

<script>
import { useGlobalConfig } from 'vuestic-ui'

export default {
	name: 'profile-section',
	data () {
		return {
		isShown: false,
		}
	},
	props: {
	options: {
		type: Array,
		default: () => [
		{
			name: 'profile',
			redirectTo: '',
			click_event: 'viewProfile'
		},
		{
			name: 'logout',
			redirectTo: 'login',
			click_event: 'logOut',
		},
	],
	},
  },
  computed: {
    theme() { return useGlobalConfig().getGlobalConfig() },
  },
  methods: {
		logOut() {
			// console.log('test')
			this.$store.dispatch('auth/logout');
			this.$router.push('/auth/login');
		},
		viewProfile() {
			return;
		}
	}
}
</script>

<style lang="scss">

.profile-dropdown {
  cursor: pointer;

  .va-dropdown-popper__anchor {
    display: flex;
    justify-content: flex-end;
  }

  &__content {
    width: 8rem;
  }

  &__item {
    display: block;
    color: var(--va-gray);

    &:hover,
    &:active {
      color: var(--va-primary);
    }
  }

  .va-dropdown__anchor {
    display: inline-block;
  }
}
</style>
