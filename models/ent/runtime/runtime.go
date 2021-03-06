// Code generated by entc, DO NOT EDIT.

package runtime

import (
	"context"
	"time"
	"wing/models/ent/auth"
	"wing/models/ent/jobhistory"
	"wing/models/ent/organization"
	"wing/models/ent/orgunit"
	"wing/models/ent/orgunitmember"
	"wing/models/ent/orgunitposition"
	"wing/models/ent/resource"
	"wing/models/ent/system"
	"wing/models/ent/user"
	"wing/models/schema"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	authFields := schema.Auth{}.Fields()
	_ = authFields
	// authDescLastAuthTime is the schema descriptor for last_auth_time field.
	authDescLastAuthTime := authFields[0].Descriptor()
	// auth.DefaultLastAuthTime holds the default value on creation for the last_auth_time field.
	auth.DefaultLastAuthTime = authDescLastAuthTime.Default.(func() time.Time)
	jobhistoryMixin := schema.JobHistory{}.Mixin()
	jobhistoryMixinHooks1 := jobhistoryMixin[1].Hooks()
	jobhistory.Hooks[0] = jobhistoryMixinHooks1[0]
	jobhistoryMixinFields2 := jobhistoryMixin[2].Fields()
	_ = jobhistoryMixinFields2
	jobhistoryFields := schema.JobHistory{}.Fields()
	_ = jobhistoryFields
	// jobhistoryDescCreateTime is the schema descriptor for create_time field.
	jobhistoryDescCreateTime := jobhistoryMixinFields2[0].Descriptor()
	// jobhistory.DefaultCreateTime holds the default value on creation for the create_time field.
	jobhistory.DefaultCreateTime = jobhistoryDescCreateTime.Default.(func() time.Time)
	orgunitMixin := schema.OrgUnit{}.Mixin()
	orgunitMixinHooks1 := orgunitMixin[1].Hooks()
	orgunitMixinHooks2 := orgunitMixin[2].Hooks()
	orgunit.Hooks[0] = orgunitMixinHooks1[0]
	orgunit.Hooks[1] = orgunitMixinHooks2[0]
	orgunitMixinFields3 := orgunitMixin[3].Fields()
	_ = orgunitMixinFields3
	orgunitMixinFields4 := orgunitMixin[4].Fields()
	_ = orgunitMixinFields4
	orgunitFields := schema.OrgUnit{}.Fields()
	_ = orgunitFields
	// orgunitDescCreateTime is the schema descriptor for create_time field.
	orgunitDescCreateTime := orgunitMixinFields3[0].Descriptor()
	// orgunit.DefaultCreateTime holds the default value on creation for the create_time field.
	orgunit.DefaultCreateTime = orgunitDescCreateTime.Default.(func() time.Time)
	// orgunitDescUpdateTime is the schema descriptor for update_time field.
	orgunitDescUpdateTime := orgunitMixinFields4[0].Descriptor()
	// orgunit.DefaultUpdateTime holds the default value on creation for the update_time field.
	orgunit.DefaultUpdateTime = orgunitDescUpdateTime.Default.(func() time.Time)
	// orgunit.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	orgunit.UpdateDefaultUpdateTime = orgunitDescUpdateTime.UpdateDefault.(func() time.Time)
	orgunitmemberMixin := schema.OrgUnitMember{}.Mixin()
	orgunitmemberMixinHooks1 := orgunitmemberMixin[1].Hooks()
	orgunitmemberMixinHooks2 := orgunitmemberMixin[2].Hooks()
	orgunitmemberHooks := schema.OrgUnitMember{}.Hooks()
	orgunitmember.Hooks[0] = orgunitmemberMixinHooks1[0]
	orgunitmember.Hooks[1] = orgunitmemberMixinHooks2[0]
	orgunitmember.Hooks[2] = orgunitmemberHooks[0]
	orgunitmember.Hooks[3] = orgunitmemberHooks[1]
	orgunitmemberMixinFields3 := orgunitmemberMixin[3].Fields()
	_ = orgunitmemberMixinFields3
	orgunitmemberMixinFields4 := orgunitmemberMixin[4].Fields()
	_ = orgunitmemberMixinFields4
	orgunitmemberFields := schema.OrgUnitMember{}.Fields()
	_ = orgunitmemberFields
	// orgunitmemberDescCreateTime is the schema descriptor for create_time field.
	orgunitmemberDescCreateTime := orgunitmemberMixinFields3[0].Descriptor()
	// orgunitmember.DefaultCreateTime holds the default value on creation for the create_time field.
	orgunitmember.DefaultCreateTime = orgunitmemberDescCreateTime.Default.(func() time.Time)
	// orgunitmemberDescUpdateTime is the schema descriptor for update_time field.
	orgunitmemberDescUpdateTime := orgunitmemberMixinFields4[0].Descriptor()
	// orgunitmember.DefaultUpdateTime holds the default value on creation for the update_time field.
	orgunitmember.DefaultUpdateTime = orgunitmemberDescUpdateTime.Default.(func() time.Time)
	// orgunitmember.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	orgunitmember.UpdateDefaultUpdateTime = orgunitmemberDescUpdateTime.UpdateDefault.(func() time.Time)
	orgunitpositionMixin := schema.OrgUnitPosition{}.Mixin()
	orgunitpositionMixinHooks1 := orgunitpositionMixin[1].Hooks()
	orgunitpositionMixinHooks2 := orgunitpositionMixin[2].Hooks()
	orgunitpositionHooks := schema.OrgUnitPosition{}.Hooks()
	orgunitposition.Hooks[0] = orgunitpositionMixinHooks1[0]
	orgunitposition.Hooks[1] = orgunitpositionMixinHooks2[0]
	orgunitposition.Hooks[2] = orgunitpositionHooks[0]
	orgunitposition.Hooks[3] = orgunitpositionHooks[1]
	orgunitpositionMixinFields3 := orgunitpositionMixin[3].Fields()
	_ = orgunitpositionMixinFields3
	orgunitpositionMixinFields4 := orgunitpositionMixin[4].Fields()
	_ = orgunitpositionMixinFields4
	orgunitpositionFields := schema.OrgUnitPosition{}.Fields()
	_ = orgunitpositionFields
	// orgunitpositionDescCreateTime is the schema descriptor for create_time field.
	orgunitpositionDescCreateTime := orgunitpositionMixinFields3[0].Descriptor()
	// orgunitposition.DefaultCreateTime holds the default value on creation for the create_time field.
	orgunitposition.DefaultCreateTime = orgunitpositionDescCreateTime.Default.(func() time.Time)
	// orgunitpositionDescUpdateTime is the schema descriptor for update_time field.
	orgunitpositionDescUpdateTime := orgunitpositionMixinFields4[0].Descriptor()
	// orgunitposition.DefaultUpdateTime holds the default value on creation for the update_time field.
	orgunitposition.DefaultUpdateTime = orgunitpositionDescUpdateTime.Default.(func() time.Time)
	// orgunitposition.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	orgunitposition.UpdateDefaultUpdateTime = orgunitpositionDescUpdateTime.UpdateDefault.(func() time.Time)
	// orgunitpositionDescLevel is the schema descriptor for level field.
	orgunitpositionDescLevel := orgunitpositionFields[2].Descriptor()
	// orgunitposition.LevelValidator is a validator for the "level" field. It is called by the builders before save.
	orgunitposition.LevelValidator = orgunitpositionDescLevel.Validators[0].(func(int) error)
	organizationMixin := schema.Organization{}.Mixin()
	organizationMixinHooks1 := organizationMixin[1].Hooks()
	organizationMixinHooks2 := organizationMixin[2].Hooks()
	organization.Hooks[0] = organizationMixinHooks1[0]
	organization.Hooks[1] = organizationMixinHooks2[0]
	organizationMixinFields3 := organizationMixin[3].Fields()
	_ = organizationMixinFields3
	organizationMixinFields4 := organizationMixin[4].Fields()
	_ = organizationMixinFields4
	organizationFields := schema.Organization{}.Fields()
	_ = organizationFields
	// organizationDescCreateTime is the schema descriptor for create_time field.
	organizationDescCreateTime := organizationMixinFields3[0].Descriptor()
	// organization.DefaultCreateTime holds the default value on creation for the create_time field.
	organization.DefaultCreateTime = organizationDescCreateTime.Default.(func() time.Time)
	// organizationDescUpdateTime is the schema descriptor for update_time field.
	organizationDescUpdateTime := organizationMixinFields4[0].Descriptor()
	// organization.DefaultUpdateTime holds the default value on creation for the update_time field.
	organization.DefaultUpdateTime = organizationDescUpdateTime.Default.(func() time.Time)
	// organization.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	organization.UpdateDefaultUpdateTime = organizationDescUpdateTime.UpdateDefault.(func() time.Time)
	resource.Policy = privacy.NewPolicies(schema.Resource{})
	resource.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := resource.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	resourceFields := schema.Resource{}.Fields()
	_ = resourceFields
	// resourceDescName is the schema descriptor for name field.
	resourceDescName := resourceFields[0].Descriptor()
	// resource.NameValidator is a validator for the "name" field. It is called by the builders before save.
	resource.NameValidator = func() func(string) error {
		validators := resourceDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// resourceDescType is the schema descriptor for type field.
	resourceDescType := resourceFields[1].Descriptor()
	// resource.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	resource.TypeValidator = func() func(string) error {
		validators := resourceDescType.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_type string) error {
			for _, fn := range fns {
				if err := fn(_type); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	systemMixin := schema.System{}.Mixin()
	systemMixinHooks1 := systemMixin[1].Hooks()
	systemMixinHooks2 := systemMixin[2].Hooks()
	system.Hooks[0] = systemMixinHooks1[0]
	system.Hooks[1] = systemMixinHooks2[0]
	systemMixinFields3 := systemMixin[3].Fields()
	_ = systemMixinFields3
	systemMixinFields4 := systemMixin[4].Fields()
	_ = systemMixinFields4
	systemFields := schema.System{}.Fields()
	_ = systemFields
	// systemDescCreateTime is the schema descriptor for create_time field.
	systemDescCreateTime := systemMixinFields3[0].Descriptor()
	// system.DefaultCreateTime holds the default value on creation for the create_time field.
	system.DefaultCreateTime = systemDescCreateTime.Default.(func() time.Time)
	// systemDescUpdateTime is the schema descriptor for update_time field.
	systemDescUpdateTime := systemMixinFields4[0].Descriptor()
	// system.DefaultUpdateTime holds the default value on creation for the update_time field.
	system.DefaultUpdateTime = systemDescUpdateTime.Default.(func() time.Time)
	// system.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	system.UpdateDefaultUpdateTime = systemDescUpdateTime.UpdateDefault.(func() time.Time)
	// systemDescName is the schema descriptor for name field.
	systemDescName := systemFields[0].Descriptor()
	// system.NameValidator is a validator for the "name" field. It is called by the builders before save.
	system.NameValidator = systemDescName.Validators[0].(func(string) error)
	userMixin := schema.User{}.Mixin()
	userMixinHooks1 := userMixin[1].Hooks()
	userMixinHooks2 := userMixin[2].Hooks()
	userHooks := schema.User{}.Hooks()
	user.Hooks[0] = userMixinHooks1[0]
	user.Hooks[1] = userMixinHooks2[0]
	user.Hooks[2] = userHooks[0]
	user.Hooks[3] = userHooks[1]
	userMixinFields3 := userMixin[3].Fields()
	_ = userMixinFields3
	userMixinFields4 := userMixin[4].Fields()
	_ = userMixinFields4
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields3[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields4[0].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	// userDescAccountName is the schema descriptor for account_name field.
	userDescAccountName := userFields[0].Descriptor()
	// user.AccountNameValidator is a validator for the "account_name" field. It is called by the builders before save.
	user.AccountNameValidator = func() func(string) error {
		validators := userDescAccountName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(account_name string) error {
			for _, fn := range fns {
				if err := fn(account_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescIsOnJob is the schema descriptor for is_on_job field.
	userDescIsOnJob := userFields[2].Descriptor()
	// user.DefaultIsOnJob holds the default value on creation for the is_on_job field.
	user.DefaultIsOnJob = userDescIsOnJob.Default.(bool)
	// userDescFamilyName is the schema descriptor for family_name field.
	userDescFamilyName := userFields[3].Descriptor()
	// user.FamilyNameValidator is a validator for the "family_name" field. It is called by the builders before save.
	user.FamilyNameValidator = func() func(string) error {
		validators := userDescFamilyName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(family_name string) error {
			for _, fn := range fns {
				if err := fn(family_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescGivenName is the schema descriptor for given_name field.
	userDescGivenName := userFields[4].Descriptor()
	// user.GivenNameValidator is a validator for the "given_name" field. It is called by the builders before save.
	user.GivenNameValidator = func() func(string) error {
		validators := userDescGivenName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(given_name string) error {
			for _, fn := range fns {
				if err := fn(given_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescDisplayName is the schema descriptor for display_name field.
	userDescDisplayName := userFields[5].Descriptor()
	// user.DisplayNameValidator is a validator for the "display_name" field. It is called by the builders before save.
	user.DisplayNameValidator = func() func(string) error {
		validators := userDescDisplayName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(display_name string) error {
			for _, fn := range fns {
				if err := fn(display_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescIDNumber is the schema descriptor for id_number field.
	userDescIDNumber := userFields[7].Descriptor()
	// user.IDNumberValidator is a validator for the "id_number" field. It is called by the builders before save.
	user.IDNumberValidator = userDescIDNumber.Validators[0].(func(string) error)
	// userDescPhoneNumber is the schema descriptor for phone_number field.
	userDescPhoneNumber := userFields[9].Descriptor()
	// user.PhoneNumberValidator is a validator for the "phone_number" field. It is called by the builders before save.
	user.PhoneNumberValidator = userDescPhoneNumber.Validators[0].(func(string) error)
	// userDescAddress is the schema descriptor for address field.
	userDescAddress := userFields[10].Descriptor()
	// user.AddressValidator is a validator for the "address" field. It is called by the builders before save.
	user.AddressValidator = userDescAddress.Validators[0].(func(string) error)
	// userDescStaffID is the schema descriptor for staff_id field.
	userDescStaffID := userFields[11].Descriptor()
	// user.StaffIDValidator is a validator for the "staff_id" field. It is called by the builders before save.
	user.StaffIDValidator = userDescStaffID.Validators[0].(func(string) error)
	// userDescPersonalEmail is the schema descriptor for personal_email field.
	userDescPersonalEmail := userFields[12].Descriptor()
	// user.PersonalEmailValidator is a validator for the "personal_email" field. It is called by the builders before save.
	user.PersonalEmailValidator = userDescPersonalEmail.Validators[0].(func(string) error)
	// userDescIntranetWorkEmail is the schema descriptor for intranet_work_email field.
	userDescIntranetWorkEmail := userFields[13].Descriptor()
	// user.IntranetWorkEmailValidator is a validator for the "intranet_work_email" field. It is called by the builders before save.
	user.IntranetWorkEmailValidator = func() func(string) error {
		validators := userDescIntranetWorkEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(intranet_work_email string) error {
			for _, fn := range fns {
				if err := fn(intranet_work_email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescExtranetWorkEmail is the schema descriptor for extranet_work_email field.
	userDescExtranetWorkEmail := userFields[14].Descriptor()
	// user.ExtranetWorkEmailValidator is a validator for the "extranet_work_email" field. It is called by the builders before save.
	user.ExtranetWorkEmailValidator = userDescExtranetWorkEmail.Validators[0].(func(string) error)
}

const (
	Version = "(devel)" // Version of ent codegen.
)
