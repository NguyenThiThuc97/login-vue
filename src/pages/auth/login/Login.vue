<template>
<div class="hold-transition login-page">
	<div class="login-box">
		<div class="login-logo">
			<a href="../../index2.html"><b>Admin</b>LTE</a>
		</div>
		<!-- /.login-logo -->
		<div class="card">
			<div class="card-body login-card-body">
				<p class="login-box-msg">Sign in to start your session</p>

				<Form @submit="handleLogin" :validation-schema="schema" class="margin-t">
					<div class="input-group mb-3">
						<Field name="username" type="text" class="form-control" placeholder="Username" autocomplete="off"/>
						<div class="input-group-append">
							<div class="input-group-text">
								<font-awesome-icon icon="envelope"/>
							</div>
						</div>
						<ErrorMessage name="username" class="invalid-feedback" style="display:block" />
					</div>
					<div class="input-group mb-3">
						<Field name="password" type="password" class="form-control" placeholder="Password" />
						<div class="input-group-append">
							<div class="input-group-text">
								<font-awesome-icon icon="lock"/>
							</div>
						</div>
						<ErrorMessage name="password"  class="invalid-feedback" style="display:block" />
					</div>
					<div class="row">
						<div class="col-8">
							<div class="icheck-primary">
								<input type="checkbox" id="remember">
								<label for="remember">
									Remember Me
								</label>
							</div>
						</div>
						<!-- /.col -->
						<div class="col-4">
							<button class="btn btn-primary btn-block" :disabled="loading">
								<span
								v-show="loading"
								class=""
								></span>
								<span>Sign In</span>
							</button>
						</div>
					<!-- /.col -->
					</div>
				</Form>
				<!-- /.social-auth-links -->

				<p class="mb-1">
					<router-link to="recover-password">I forgot my password</router-link>
				</p>
				<p class="mb-0">
					<router-link to="/auth/signup" class="text-center">Register a new membership</router-link>
				</p>
				<div class="form-group">
					<div v-if="message" class="alert alert-danger" role="alert">
						{{ message }}
					</div>
				</div>
			</div>
			<!-- /.login-card-body -->
		</div>
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
			loggedIn() {
				return this.$store.state.auth.status.loggedIn;
			},
		},
		created() {
			if (this.loggedIn) {
				this.$router.push("/home");
			}
		},
		methods: {
			handleLogin(user) {
				this.loading = true;

				this.$store.dispatch("auth/login", user, {root:true}).then(
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