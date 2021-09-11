export declare module Discord {
  export interface User {
    id: string;
    email: string;
    username: string;
    avatar: string;
    locale: string;
    discriminator: string;
    token: string;
    verified: boolean;
    mfa_enabled: boolean;
    bot: boolean;
    public_flags: number;
    premium_type: number;
    system: boolean;
    flags: number;
  }

  module Guilds {
    export interface Role {
      id: string;
      name: string;
      managed: boolean;
      mentionable: boolean;
      hoist: boolean;
      color: number;
      position: number;
      permissions: string;
    }

    export interface Emoji {
      id: string;
      name: string;
      roles: any[];
      user?: any;
      require_colons: boolean;
      managed: boolean;
      animated: boolean;
      available: boolean;
    }

    export interface User {
      id: string;
      email: string;
      username: string;
      avatar: string;
      locale: string;
      discriminator: string;
      token: string;
      verified: boolean;
      mfa_enabled: boolean;
      bot: boolean;
      public_flags: number;
      premium_type: number;
      system: boolean;
      flags: number;
    }

    export interface Member {
      guild_id: string;
      joined_at: Date;
      nick: string;
      deaf: boolean;
      mute: boolean;
      user: User;
      roles: string[];
      premium_since: string;
      pending: boolean;
    }

    export interface PermissionOverwrite {
      id: string;
      type: number;
      deny: string;
      allow: string;
    }

    export interface Channel {
      id: string;
      guild_id: string;
      name: string;
      topic: string;
      type: number;
      last_message_id: string;
      last_pin_timestamp: any;
      nsfw: boolean;
      icon: string;
      position: number;
      bitrate: number;
      recipients?: any;
      permission_overwrites: PermissionOverwrite[];
      user_limit: number;
      parent_id: string;
      rate_limit_per_user: number;
      owner_id: string;
      application_id: string;
    }

    export interface Guild {
      id: string;
      name: string;
      icon: string;
      region: string;
      afk_channel_id: string;
      owner_id: string;
      owner: boolean;
      joined_at: Date;
      discovery_splash: string;
      splash: string;
      afk_timeout: number;
      member_count: number;
      verification_level: number;
      large: boolean;
      default_message_notifications: number;
      roles: Role[];
      emojis: Emoji[];
      members: Member[];
      presences: any[];
      max_presences: number;
      max_members: number;
      channels: Channel[];
      voice_states: any[];
      unavailable: boolean;
      explicit_content_filter: number;
      features: any[];
      mfa_level: number;
      application_id: string;
      widget_enabled: boolean;
      widget_channel_id: string;
      system_channel_id: string;
      system_channel_flags: number;
      rules_channel_id: string;
      vanity_url_code: string;
      description: string;
      banner: string;
      premium_tier: number;
      premium_subscription_count: number;
      preferred_locale: string;
      public_updates_channel_id: string;
      max_video_channel_users: number;
      approximate_member_count: number;
      approximate_presence_count: number;
      permissions: string;
    }
  }
}
