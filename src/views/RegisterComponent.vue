<template>
	<div class="login-container text-c animated flipInX">
		<div>
			<h1 class="logo-badge text-whitesmoke"><span class="fa fa-user-circle"></span></h1>
		</div>
			<h3 class="text-whitesmoke">Sign Up</h3>
			<!-- <p class="text-whitesmoke">Sign In</p> -->
		<div class="container-content">
			<Form @submit="handleRegister" :validation-schema="schema" class="margin-t">
				<div class="form-group">
					<Field name="username" type="text" class="form-control" placeholder="Username" />
					<ErrorMessage name="username" class="error-feedback" />
				</div>
				<div class="form-group">
					<Field name="email" type="text" class="form-control" placeholder="Email" />
					<ErrorMessage name="email" class="error-feedback" />
				</div>
				<div class="form-group">
					<Field name="password" type="password" class="form-control" placeholder="Password" />
					<ErrorMessage name="password" class="error-feedback" />
				</div>
				<div class="form-group">
					<Field name="re_password" type="password" class="form-control" placeholder="Confirm Password" />
					<ErrorMessage name="re_password" class="error-feedback" />
				</div>
				<button class="form-button button-l margin-b" :disabled="loading">
					<span
					v-show="loading"
					class="spinner-border spinner-border-sm"
					></span>
					<span>Sign Up</span>
				</button>
				<p class="text-whitesmoke text-center"><small>Have already an account?</small></p>
				<router-link to="/signin" class="text-darkyellow"><small>Sign In</small></router-link>
				<div class="form-group">
					<div v-if="message" class="alert alert-danger" role="alert">
						{{ message }}
					</div>
				</div>
			</form>
			<p class="margin-t text-whitesmoke"><small> concobebe &copy; 2021</small> </p>
		</div>
	</div>
</template>

<script>
	import { Form, Field, ErrorMessage } from 'vee-validate';
	import * as yup from "yup";

	export default {
		name: "Login",
		components: {
			Form,
			Field,
			ErrorMessage,
		},
		data() {
			const schema = yup.object().shape({
				username: yup.string().required("Username is required!"),
				password: yup.string().required("Password is required!"),
			});

			return {
				loading: false,
				message: "",
				schema,
			};
		},
		computed: {
			// loggedIn() {
			// return this.$store.state.auth.status.loggedIn;
			// },
		},
		created() {
			// if (this.loggedIn) {
			// this.$router.push("/profile");
			// }
		},
		methods: {
			handleRegister(user) {
				this.loading = true;

				this.$store.dispatch("auth/register", user, {root:true}).then(
					() => {
						this.$router.push("/");
					},
					(error) => {
						this.loading = false;
						this.message =
							(error.response &&
							error.response.data &&
							error.response.data.message) ||
							error.message ||
							error.toString();
					}
				);
			},
		},
	};
</script>