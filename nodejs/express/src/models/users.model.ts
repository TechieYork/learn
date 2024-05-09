type UserProfileOptions = (u: UserProfile) => void;

class UserProfile {
  // Define the properties of a user profile here
  public id: string;
  public name: string;

  constructor(...options: UserProfileOptions[]) {
    this.id = "";
    this.name = "";

    // set the options
    for (const option of options) {
      option(this);
    }
  }

  public static WithID(id: string) {
    return (u: UserProfile): void => {
      u.id = id;
    };
  }

  public static WithName(name: string) {
    return (u: UserProfile): void => {
      u.name = name;
    };
  }
}

class UserStorage {
  async get(id: string): Promise<UserProfile | undefined> {
    // to be implemented, for now return a dummy user profile
    // for instance, send query to mongoDB or other database
    return new UserProfile(
      UserProfile.WithID("test-id"),
      UserProfile.WithName("test-name")
    );
  }
}

export { UserProfile, UserProfileOptions, UserStorage };
