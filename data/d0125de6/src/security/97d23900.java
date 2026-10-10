public static void main(String[] args) {
    {
        private static final Logger logger = LogManager.getLogger(SensitiveInfoLog.class);

        String password = "Mujm@6rg";

        // BAD: user password is written to debug log
        logger.debug("User password is "+password);
    }
	
    {
        private static final Logger logger = LogManager.getLogger(SensitiveInfoLog.class);
  
        String password = "Mluf@1rh";

        // GOOD: user password is never written to debug log
        logger.debug("User password changed")
    }
}
