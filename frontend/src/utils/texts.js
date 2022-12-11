const language = localStorage.getItem('language') || 'en';

// Tables
const HighPriorityTasksText = language === 'en' ? 'Highest Priority Tasks' : 'Tareas de alta prioridad';
const TotalNumberOfTasksText = language === 'en' ? 'Total number of tasks' : 'Número total de tareas';
const HeaderTasksText = language === 'en' ? 'Tasks' : 'Tareas';
const HeaderCategoryText = language === 'en' ? 'Category' : 'Categoría';
const HeaderActionsText = language === 'en' ? 'Actions' : 'Acciones';
const HeaderNameText = language === 'en' ? 'Name' : 'Nombre';
const CompletedTodosIndicationText = language === 'en' ? 'You better complete some Todos first!!!' : '¡¡¡Debes completar alguna tarea antes!!!';
const RecurringTodosIndicationText = language === 'en' ? 'You better create some recurring Todos first!!!' : '¡¡¡Debes crear alguna tarea recurrente antes!!!';

// Icons
const ShareIconText = language === 'en' ? 'Share' : 'Compartir';
const UnSubscribeIconText = language === 'en' ? 'Unsubscribe' : 'Darse de baja';
const DeleteIconText = language === 'en' ? 'Delete' : 'Eliminar';
const EditIconText = language === 'en' ? 'Edit' : 'Editar';
const ViewIconText = language === 'en' ? 'View' : 'Ver';
const ReactivateIconText = language === 'en' ? 'Reactivate' : 'Reactivar';
const RecurringTodosIconText = language === 'en' ? 'Recurring Todos' : 'Tareas recurrentes';
const StatisticsIconText = language === 'en' ? 'Statistics (coming soon)' : 'Estadísticas (próximamente)';
const ReportABugIconText = language === 'en' ? 'Report a bug' : 'Reportar un error';
const ConfigurationIconText = language === 'en' ? 'Configuration (coming soon)' : 'Configuración (próximamente)';
const LogoutIconText = language === 'en' ? 'Logout' : 'Cerrar sesión';
const StartIconText = language === 'en' ? 'Start' : 'Iniciar';
const CompleteIconText = language === 'en' ? 'Complete' : 'Completar';
const CreateIconText = language === 'en' ? 'Create a new Todo' : 'Crear una nueva tarea';
const CreateCategoryIconText = language === 'en' ? 'Create a new Category' : 'Crear una nueva categoría';
const CategoriesIconText = language === 'en' ? 'Categories' : 'Categorías';
const CompletedTodosIconText = language === 'en' ? 'Completed Todos' : 'Tareas completadas';
const ProfileIconText = language === 'en' ? 'Profile' : 'Perfil';

// Modals
const UserAlreadySubscribedText = language === 'en' ? 'User already subscribed to that category!' : '¡No se puede suscribir a una categoría ya suscrita!';
const OnlyOwnersCanDeleteCategoryText = language === 'en' ? 'Only owners can delete a shared category. If you want the category to disappear, unsubscribe from it!' : '¡Solo los propietarios pueden eliminar una categoría compartida. Si quieres que desaparezca, date de baja de la misma';
const OnlyOwnerCanEditCategoryText = language === 'en' ? 'Only owners can edit a shared category!' : '¡Solo los propietarios pueden editar una categoría compartida!';
const ConfirmUnshareCategoryText = language === 'en' ? 'Are you sure you want to unsubscribe from this category?' : '¿Estás seguro de que quieres darte de baja de esta categoría?';
const CategoryAlreadyExistsText = language === 'en' ? 'Category already exists! Please try with a different name' : '¡La categoría ya existe! Por favor, inténtelo con un nombre diferente';
const UserNotFoundText = language === 'en' ? 'User not found! Please try with a different username or register first' : '¡Usuario no encontrado! Por favor, inténtelo con un nombre de usuario diferente o regístrese primero';
const EnterEmailText = language === 'en' ? 'Please enter your email' : 'Por favor, introduzca su correo electrónico';
const PasswordNotLongEnoughText = language === 'en' ? 'Password must be at least 13 characters long' : 'La contraseña debe tener al menos 13 caracteres';
const IncorrectPasswordText = language === 'en' ? 'Incorrect password! Please try again' : '¡Contraseña incorrecta! Por favor, inténtelo de nuevo';
const PasswordsDoNotMatchText = language === 'en' ? 'Passwords do not match! Please try again' : '¡Las contraseñas no coinciden! Por favor, inténtelo de nuevo';
const UserAlreadyRegisteredText = language === 'en' ? 'User already registered! Please try with a different username' : '¡Usuario ya registrado! Por favor, inténtelo con un nombre de usuario diferente';
const InvalidEmailText = language === 'en' ? 'Invalid email address format! Please try again with a valid format' : '¡Formato de dirección de correo electrónico no válido! Por favor, inténtelo de nuevo con un formato válido';
const DeletingCategoryText = language === 'en' ? 'Warning! Deleting a category will delete all its Todos. Are you sure you want to delete this category?' : '¡Advertencia! Eliminar una categoría eliminará todas sus tareas. ¿Estás segur@ de que quieres eliminar esta categoría?';

// Buttons
const CancelButtonText = language === 'en' ? 'Cancel' : 'Cancelar';
const ShareButtonText = language === 'en' ? 'Share' : 'Compartir';
const ReturnButtonText = language === 'en' ? 'Return' : 'Volver';
const CreateButtonText = language === 'en' ? 'Create' : 'Crear';
const LoginButtonText = language === 'en' ? 'Login' : 'Iniciar sesión';
const RegisterButtonText = language === 'en' ? 'Register' : 'Registrarse';
const ReportButtonText = language === 'en' ? 'Report' : 'Reportar';
const UnsuscribeButtonText = language === 'en' ? 'Unsubscribe' : 'Darse de baja';
const EditButtonText = language === 'en' ? 'Edit' : 'Editar';
const DeleteButtonText = language === 'en' ? 'Delete' : 'Eliminar';

// Headers
const CategoriesHeaderText = language === 'en' ? 'Categories' : 'Categorías';
const LoginHeaderText = language === 'en' ? 'Login' : 'Iniciar sesión';
const EditCategoryHeaderText = language === 'en' ? 'Edit category' : 'Editar categoría';
const ShareCategoryHeaderText = language === 'en' ? 'Share category' : 'Compartir categoría';
const ViewCategoryHeaderText = language === 'en' ? 'View category' : 'Ver categoría';
const CompletedTodosHeaderText = language === 'en' ? 'Completed Todos' : 'Tareas completadas';
const CreateCategoryHeaderText = language === 'en' ? 'Create category' : 'Crear categoría';
const CreateTodoHeaderText = language === 'en' ? 'Create Todo' : 'Crear tarea';
const ThankYouHeaderText = language === 'en' ? 'Thank you for using DAPS!' : '¡Gracias por usar DAPS!';
const RecurringTodosHeaderText = language === 'en' ? 'Recurring Todos' : 'Tareas recurrentes';
const RegisterHeaderText = language === 'en' ? 'Register' : 'Registrarse';
const ReportABugHeaderText = language === 'en' ? 'Report a bug' : 'Reportar un error';
const EditTodoHeaderText = language === 'en' ? 'Edit Todo' : 'Editar tarea';
const ViewTodoHeaderText = language === 'en' ? 'View Todo' : 'Ver tarea';
const WelcomeToDapsText = language === 'en' ? 'Welcome to DAPS!' : '¡Bienvenid@ a DAPS!';
const ProfileHeaderText = language === 'en' ? 'Profile' : 'Perfil';


// Forms
const NameLabelText = language === 'en' ? 'Name' : 'Nombre';
const DescriptionLabelText = language === 'en' ? 'Description' : 'Descripción';
const PriorityLabelText = language === 'en' ? 'Priority' : 'Prioridad';
const SelectPriorityText = language === 'en' ? 'Select priority' : 'Selecciona prioridad';
const LowestPriorityText = language === 'en' ? 'Lowest' : 'Muy baja';
const LowPriorityText = language === 'en' ? 'Low' : 'Baja';
const MediumPriorityText = language === 'en' ? 'Medium' : 'Media';
const HighPriorityText = language === 'en' ? 'High' : 'Alta';
const HighestPriorityText = language === 'en' ? 'Highest' : 'Máxima';
const RecurringLabelText = language === 'en' ? 'Recurrency' : 'Recurrente';
const SelectRecurringText = language === 'en' ? 'Select recurring' : 'Selecciona recurrencia';
const YesRecurringText = language === 'en' ? 'Yes' : 'Sí';
const NoRecurringText = language === 'en' ? 'No' : 'No';
const LinkLabelText = language === 'en' ? 'Link' : 'Enlace';
const EmailAddressLabelText = language === 'en' ? 'Email address' : 'Correo electrónico';
const PasswordLabelText = language === 'en' ? 'Password' : 'Contraseña';
const RepeatPasswordLabelText = language === 'en' ? 'Repeat password' : 'Repita contraseña';
const LanguageLabelText = language === 'en' ? 'Language' : 'Idioma';
const EnglishLanguageText = language === 'en' ? 'English' : 'Inglés';
const SpanishLanguageText = language === 'en' ? 'Spanish' : 'Español';
const AutoSuggestLabelText = language === 'en' ? 'Auto-suggest' : 'Auto-sugerencia';

export {
    HighPriorityTasksText,
    TotalNumberOfTasksText,
    HeaderTasksText,
    HeaderCategoryText,
    HeaderActionsText,
    CompletedTodosIndicationText,
    ShareIconText,
    DeleteIconText,
    EditIconText,
    ViewIconText,
    ReactivateIconText,
    RecurringTodosIconText,
    StatisticsIconText,
    ReportABugIconText,
    ConfigurationIconText,
    LogoutIconText,
    StartIconText,
    CompleteIconText,
    CreateIconText,
    CreateCategoryIconText,
    UserAlreadySubscribedText,
    OnlyOwnersCanDeleteCategoryText,
    OnlyOwnerCanEditCategoryText,
    ConfirmUnshareCategoryText,
    CategoryAlreadyExistsText,
    UserNotFoundText,
    EnterEmailText,
    PasswordNotLongEnoughText,
    IncorrectPasswordText,
    PasswordsDoNotMatchText,
    UserAlreadyRegisteredText,
    CancelButtonText,
    ShareButtonText,
    ReturnButtonText,
    CreateButtonText,
    LoginButtonText,
    RegisterButtonText,
    ReportButtonText,
    CategoriesHeaderText,
    EditCategoryHeaderText,
    ViewCategoryHeaderText,
    CompletedTodosHeaderText,
    CreateCategoryHeaderText,
    CreateTodoHeaderText,
    ThankYouHeaderText,
    RecurringTodosHeaderText,
    RegisterHeaderText,
    NameLabelText,
    DescriptionLabelText,
    PriorityLabelText,
    SelectPriorityText,
    LowestPriorityText,
    LowPriorityText,
    MediumPriorityText,
    HighPriorityText,
    HighestPriorityText,
    RecurringLabelText,
    YesRecurringText,
    NoRecurringText,
    LinkLabelText,
    EmailAddressLabelText,
    PasswordLabelText,
    RepeatPasswordLabelText,
    LoginHeaderText,
    UnSubscribeIconText,
    ShareCategoryHeaderText,
    UnsuscribeButtonText,
    EditButtonText,
    SelectRecurringText,
    CategoriesIconText,
    CompletedTodosIconText,
    HeaderNameText,
    ReportABugHeaderText,
    EditTodoHeaderText,
    ViewTodoHeaderText,
    RecurringTodosIndicationText,
    WelcomeToDapsText,
    InvalidEmailText,
    ProfileHeaderText,
    LanguageLabelText,
    AutoSuggestLabelText,
    EnglishLanguageText,
    SpanishLanguageText,
    ProfileIconText,
    DeletingCategoryText,
    DeleteButtonText,
};