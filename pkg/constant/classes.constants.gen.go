package constant

/*------------------------------------------------------------------------------
//   This code was generated by template classes.constants.go.tmpl.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "classes.constants.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//revive:disable

const (
	ANIMATED_TEXTURE_MAX_FRAMES int32 = 256
)

const (
	ANIMATION_NODE_BLEND_TREE_CONNECTION_OK                      int32 = 0
	ANIMATION_NODE_BLEND_TREE_CONNECTION_ERROR_NO_INPUT          int32 = 1
	ANIMATION_NODE_BLEND_TREE_CONNECTION_ERROR_NO_INPUT_INDEX    int32 = 2
	ANIMATION_NODE_BLEND_TREE_CONNECTION_ERROR_NO_OUTPUT         int32 = 3
	ANIMATION_NODE_BLEND_TREE_CONNECTION_ERROR_SAME_NODE         int32 = 4
	ANIMATION_NODE_BLEND_TREE_CONNECTION_ERROR_CONNECTION_EXISTS int32 = 5
)

const (
	AUDIO_STREAM_PLAYBACK_POLYPHONIC_INVALID_ID int32 = -1
)

const (
	CANVAS_ITEM_NOTIFICATION_TRANSFORM_CHANGED       int32 = 2000
	CANVAS_ITEM_NOTIFICATION_LOCAL_TRANSFORM_CHANGED int32 = 35
	CANVAS_ITEM_NOTIFICATION_DRAW                    int32 = 30
	CANVAS_ITEM_NOTIFICATION_VISIBILITY_CHANGED      int32 = 31
	CANVAS_ITEM_NOTIFICATION_ENTER_CANVAS            int32 = 32
	CANVAS_ITEM_NOTIFICATION_EXIT_CANVAS             int32 = 33
	CANVAS_ITEM_NOTIFICATION_WORLD_2_D_CHANGED       int32 = 36
)

const (
	CONTAINER_NOTIFICATION_PRE_SORT_CHILDREN int32 = 50
	CONTAINER_NOTIFICATION_SORT_CHILDREN     int32 = 51
)

const (
	CONTROL_NOTIFICATION_RESIZED                  int32 = 40
	CONTROL_NOTIFICATION_MOUSE_ENTER              int32 = 41
	CONTROL_NOTIFICATION_MOUSE_EXIT               int32 = 42
	CONTROL_NOTIFICATION_MOUSE_ENTER_SELF         int32 = 60
	CONTROL_NOTIFICATION_MOUSE_EXIT_SELF          int32 = 61
	CONTROL_NOTIFICATION_FOCUS_ENTER              int32 = 43
	CONTROL_NOTIFICATION_FOCUS_EXIT               int32 = 44
	CONTROL_NOTIFICATION_THEME_CHANGED            int32 = 45
	CONTROL_NOTIFICATION_SCROLL_BEGIN             int32 = 47
	CONTROL_NOTIFICATION_SCROLL_END               int32 = 48
	CONTROL_NOTIFICATION_LAYOUT_DIRECTION_CHANGED int32 = 49
)

const (
	DISPLAY_SERVER_SCREEN_WITH_MOUSE_FOCUS    int32 = -4
	DISPLAY_SERVER_SCREEN_WITH_KEYBOARD_FOCUS int32 = -3
	DISPLAY_SERVER_SCREEN_PRIMARY             int32 = -2
	DISPLAY_SERVER_SCREEN_OF_MAIN_WINDOW      int32 = -1
	DISPLAY_SERVER_MAIN_WINDOW_ID             int32 = 0
	DISPLAY_SERVER_INVALID_WINDOW_ID          int32 = -1
	DISPLAY_SERVER_INVALID_INDICATOR_ID       int32 = -1
)

const (
	E_NET_PACKET_PEER_PACKET_LOSS_SCALE        int32 = 65536
	E_NET_PACKET_PEER_PACKET_THROTTLE_SCALE    int32 = 32
	E_NET_PACKET_PEER_FLAG_RELIABLE            int32 = 1
	E_NET_PACKET_PEER_FLAG_UNSEQUENCED         int32 = 2
	E_NET_PACKET_PEER_FLAG_UNRELIABLE_FRAGMENT int32 = 8
)

const (
	EDITOR_SCENE_FORMAT_IMPORTER_IMPORT_SCENE                          int32 = 1
	EDITOR_SCENE_FORMAT_IMPORTER_IMPORT_ANIMATION                      int32 = 2
	EDITOR_SCENE_FORMAT_IMPORTER_IMPORT_FAIL_ON_MISSING_DEPENDENCIES   int32 = 4
	EDITOR_SCENE_FORMAT_IMPORTER_IMPORT_GENERATE_TANGENT_ARRAYS        int32 = 8
	EDITOR_SCENE_FORMAT_IMPORTER_IMPORT_USE_NAMED_SKIN_BINDS           int32 = 16
	EDITOR_SCENE_FORMAT_IMPORTER_IMPORT_DISCARD_MESHES_AND_MATERIALS   int32 = 32
	EDITOR_SCENE_FORMAT_IMPORTER_IMPORT_FORCE_DISABLE_MESH_COMPRESSION int32 = 64
)

const (
	EDITOR_SETTINGS_NOTIFICATION_EDITOR_SETTINGS_CHANGED int32 = 10000
)

const (
	GLTF_STATE_HANDLE_BINARY_DISCARD_TEXTURES      int32 = 0
	GLTF_STATE_HANDLE_BINARY_EXTRACT_TEXTURES      int32 = 1
	GLTF_STATE_HANDLE_BINARY_EMBED_AS_BASISU       int32 = 2
	GLTF_STATE_HANDLE_BINARY_EMBED_AS_UNCOMPRESSED int32 = 3
)

const (
	GPU_PARTICLES_3_D_MAX_DRAW_PASSES int32 = 4
)

const (
	GRID_MAP_INVALID_CELL_ITEM int32 = -1
)

const (
	IP_RESOLVER_MAX_QUERIES int32 = 256
	IP_RESOLVER_INVALID_ID  int32 = -1
)

const (
	IMAGE_MAX_WIDTH  int32 = 16777216
	IMAGE_MAX_HEIGHT int32 = 16777216
)

const (
	INPUT_EVENT_DEVICE_ID_EMULATION int32 = -1
)

const (
	MAIN_LOOP_NOTIFICATION_OS_MEMORY_WARNING     int32 = 2009
	MAIN_LOOP_NOTIFICATION_TRANSLATION_CHANGED   int32 = 2010
	MAIN_LOOP_NOTIFICATION_WM_ABOUT              int32 = 2011
	MAIN_LOOP_NOTIFICATION_CRASH                 int32 = 2012
	MAIN_LOOP_NOTIFICATION_OS_IME_UPDATE         int32 = 2013
	MAIN_LOOP_NOTIFICATION_APPLICATION_RESUMED   int32 = 2014
	MAIN_LOOP_NOTIFICATION_APPLICATION_PAUSED    int32 = 2015
	MAIN_LOOP_NOTIFICATION_APPLICATION_FOCUS_IN  int32 = 2016
	MAIN_LOOP_NOTIFICATION_APPLICATION_FOCUS_OUT int32 = 2017
	MAIN_LOOP_NOTIFICATION_TEXT_SERVER_CHANGED   int32 = 2018
)

const (
	MATERIAL_RENDER_PRIORITY_MAX int32 = 127
	MATERIAL_RENDER_PRIORITY_MIN int32 = -128
)

const (
	MULTIPLAYER_PEER_TARGET_PEER_BROADCAST int32 = 0
	MULTIPLAYER_PEER_TARGET_PEER_SERVER    int32 = 1
)

const (
	NODE_NOTIFICATION_ENTER_TREE                  int32 = 10
	NODE_NOTIFICATION_EXIT_TREE                   int32 = 11
	NODE_NOTIFICATION_MOVED_IN_PARENT             int32 = 12
	NODE_NOTIFICATION_READY                       int32 = 13
	NODE_NOTIFICATION_PAUSED                      int32 = 14
	NODE_NOTIFICATION_UNPAUSED                    int32 = 15
	NODE_NOTIFICATION_PHYSICS_PROCESS             int32 = 16
	NODE_NOTIFICATION_PROCESS                     int32 = 17
	NODE_NOTIFICATION_PARENTED                    int32 = 18
	NODE_NOTIFICATION_UNPARENTED                  int32 = 19
	NODE_NOTIFICATION_SCENE_INSTANTIATED          int32 = 20
	NODE_NOTIFICATION_DRAG_BEGIN                  int32 = 21
	NODE_NOTIFICATION_DRAG_END                    int32 = 22
	NODE_NOTIFICATION_PATH_RENAMED                int32 = 23
	NODE_NOTIFICATION_CHILD_ORDER_CHANGED         int32 = 24
	NODE_NOTIFICATION_INTERNAL_PROCESS            int32 = 25
	NODE_NOTIFICATION_INTERNAL_PHYSICS_PROCESS    int32 = 26
	NODE_NOTIFICATION_POST_ENTER_TREE             int32 = 27
	NODE_NOTIFICATION_DISABLED                    int32 = 28
	NODE_NOTIFICATION_ENABLED                     int32 = 29
	NODE_NOTIFICATION_RESET_PHYSICS_INTERPOLATION int32 = 2001
	NODE_NOTIFICATION_EDITOR_PRE_SAVE             int32 = 9001
	NODE_NOTIFICATION_EDITOR_POST_SAVE            int32 = 9002
	NODE_NOTIFICATION_WM_MOUSE_ENTER              int32 = 1002
	NODE_NOTIFICATION_WM_MOUSE_EXIT               int32 = 1003
	NODE_NOTIFICATION_WM_WINDOW_FOCUS_IN          int32 = 1004
	NODE_NOTIFICATION_WM_WINDOW_FOCUS_OUT         int32 = 1005
	NODE_NOTIFICATION_WM_CLOSE_REQUEST            int32 = 1006
	NODE_NOTIFICATION_WM_GO_BACK_REQUEST          int32 = 1007
	NODE_NOTIFICATION_WM_SIZE_CHANGED             int32 = 1008
	NODE_NOTIFICATION_WM_DPI_CHANGE               int32 = 1009
	NODE_NOTIFICATION_VP_MOUSE_ENTER              int32 = 1010
	NODE_NOTIFICATION_VP_MOUSE_EXIT               int32 = 1011
	NODE_NOTIFICATION_OS_MEMORY_WARNING           int32 = 2009
	NODE_NOTIFICATION_TRANSLATION_CHANGED         int32 = 2010
	NODE_NOTIFICATION_WM_ABOUT                    int32 = 2011
	NODE_NOTIFICATION_CRASH                       int32 = 2012
	NODE_NOTIFICATION_OS_IME_UPDATE               int32 = 2013
	NODE_NOTIFICATION_APPLICATION_RESUMED         int32 = 2014
	NODE_NOTIFICATION_APPLICATION_PAUSED          int32 = 2015
	NODE_NOTIFICATION_APPLICATION_FOCUS_IN        int32 = 2016
	NODE_NOTIFICATION_APPLICATION_FOCUS_OUT       int32 = 2017
	NODE_NOTIFICATION_TEXT_SERVER_CHANGED         int32 = 2018
)

const (
	NODE_3_D_NOTIFICATION_TRANSFORM_CHANGED       int32 = 2000
	NODE_3_D_NOTIFICATION_ENTER_WORLD             int32 = 41
	NODE_3_D_NOTIFICATION_EXIT_WORLD              int32 = 42
	NODE_3_D_NOTIFICATION_VISIBILITY_CHANGED      int32 = 43
	NODE_3_D_NOTIFICATION_LOCAL_TRANSFORM_CHANGED int32 = 44
)

const (
	OBJECT_NOTIFICATION_POSTINITIALIZE     int32 = 0
	OBJECT_NOTIFICATION_PREDELETE          int32 = 1
	OBJECT_NOTIFICATION_EXTENSION_RELOADED int32 = 2
)

const (
	RD_FRAMEBUFFER_PASS_ATTACHMENT_UNUSED int32 = -1
)

const (
	RENDERING_DEVICE_INVALID_ID        int32 = -1
	RENDERING_DEVICE_INVALID_FORMAT_ID int32 = -1
)

const (
	RENDERING_SERVER_NO_INDEX_ARRAY                     int32 = -1
	RENDERING_SERVER_ARRAY_WEIGHTS_SIZE                 int32 = 4
	RENDERING_SERVER_CANVAS_ITEM_Z_MIN                  int32 = -4096
	RENDERING_SERVER_CANVAS_ITEM_Z_MAX                  int32 = 4096
	RENDERING_SERVER_MAX_GLOW_LEVELS                    int32 = 7
	RENDERING_SERVER_MAX_CURSORS                        int32 = 8
	RENDERING_SERVER_MAX_2_D_DIRECTIONAL_LIGHTS         int32 = 8
	RENDERING_SERVER_MAX_MESH_SURFACES                  int32 = 256
	RENDERING_SERVER_MATERIAL_RENDER_PRIORITY_MIN       int32 = -128
	RENDERING_SERVER_MATERIAL_RENDER_PRIORITY_MAX       int32 = 127
	RENDERING_SERVER_ARRAY_CUSTOM_COUNT                 int32 = 4
	RENDERING_SERVER_PARTICLES_EMIT_FLAG_POSITION       int32 = 1
	RENDERING_SERVER_PARTICLES_EMIT_FLAG_ROTATION_SCALE int32 = 2
	RENDERING_SERVER_PARTICLES_EMIT_FLAG_VELOCITY       int32 = 4
	RENDERING_SERVER_PARTICLES_EMIT_FLAG_COLOR          int32 = 8
	RENDERING_SERVER_PARTICLES_EMIT_FLAG_CUSTOM         int32 = 16
)

const (
	RESOURCE_UID_INVALID_ID int32 = -1
)

const (
	SKELETON_3_D_NOTIFICATION_UPDATE_SKELETON int32 = 50
)

const (
	TILE_SET_ATLAS_SOURCE_TRANSFORM_FLIP_H    int32 = 4096
	TILE_SET_ATLAS_SOURCE_TRANSFORM_FLIP_V    int32 = 8192
	TILE_SET_ATLAS_SOURCE_TRANSFORM_TRANSPOSE int32 = 16384
)

const (
	VISUAL_SHADER_NODE_ID_INVALID int32 = -1
	VISUAL_SHADER_NODE_ID_OUTPUT  int32 = 0
)

const (
	WINDOW_NOTIFICATION_VISIBILITY_CHANGED int32 = 30
	WINDOW_NOTIFICATION_THEME_CHANGED      int32 = 32
)
