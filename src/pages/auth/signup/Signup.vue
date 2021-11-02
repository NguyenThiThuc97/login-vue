<template>
<div class="hold-transition register-page">
	<div class="register-box">
		<div class="register-logo">
			<a href="../../index2.html"><b>Admin</b>LTE</a>
		</div>

		<div class="card">
			<div class="card-body register-card-body">
			<p class="login-box-msg">Register a new membership</p>

			<Form @submit="handleRegister" :validation-schema="schema">
				<div class="input-group mb-3">
					<Field name="username" type="text" class="form-control" placeholder="Username" autocomplete="off" />
					<div class="input-group-append">
						<div class="input-group-text">
							<font-awesome-icon icon="user"/>
						</div>
					</div>
					<ErrorMessage name="username" class="invalid-feedback" style="display:block" />
				</div>
				<div class="input-group mb-3">
					<Field name="email" type="text" class="form-control" placeholder="Email" autocomplete="off" />
					<div class="input-group-append">
						<div class="input-group-text">
							<font-awesome-icon icon="envelope"/>
						</div>
					</div>
					<ErrorMessage name="email" class="invalid-feedback" style="display:block" />
				</div>
				<div class="input-group mb-3">
					<Field name="password" type="password" class="form-control" placeholder="Password" />
					<div class="input-group-append">
						<div class="input-group-text">
							<font-awesome-icon icon="lock"/>
						</div>
					</div>
					<ErrorMessage name="password" class="invalid-feedback" style="display:block" />
				</div>
				<div class="input-group mb-3">
					<Field name="re_password" type="password" class="form-control" placeholder="Retype Password" />
					<div class="input-group-append">
						<div class="input-group-text">
							<font-awesome-icon icon="lock"/>
						</div>
					</div>
					<ErrorMessage name="re_password" class="invalid-feedback" style="display:block" />
				</div>
				<div class="row">
					<div class="col-8">
						<div class="icheck-primary">
						<input type="checkbox" id="agreeTerms" name="terms" value="agree">
						<label for="agreeTerms">
						I agree to the <a href="#">terms</a>
						</label>
						</div>
					</div>
					<!-- /.col -->
					<div class="col-4">
						<button class="btn btn-primary btn-block" :disabled="loading">
							<span
							v-show="loading"
							class="spinner-border spinner-border-sm"
							></span>
							<span>Register</span>
						</button>
					</div>
					<!-- /.col -->
				</div>
			</Form>
			<router-link to="/auth/login" class="text-center">I already have a membership</router-link>
			<div class="form-group">
				<div v-if="message" class="alert alert-danger" role="alert">
					{{ message }}
				</div>
			</div>
			</div>
			<!-- /.form-box -->
		</div><!-- /.card -->
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
				email: yup.string()
					.required("Username is required!")
					.email("Email is invalid"),
				re_password: yup.string()
					.oneOf([yup.ref('password'), null], 'Passwords must match')
					.required('Confirm Password is required')
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
						this.$router.push("/admin");
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